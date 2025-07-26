class FixSchemaToMatchDesign < ActiveRecord::Migration[8.0]
  def up
    # Enable UUID extension
    enable_extension 'pgcrypto' unless extension_enabled?('pgcrypto')
    
    # First, handle audit_logs modifications (remove the products foreign key)
    if column_exists?(:audit_logs, :product_id)
      remove_foreign_key :audit_logs, :products if foreign_key_exists?(:audit_logs, :products)
      remove_column :audit_logs, :product_id
    end
    
    # Add item_id to audit_logs
    add_column :audit_logs, :item_id, :bigint unless column_exists?(:audit_logs, :item_id)
    add_foreign_key :audit_logs, :items unless foreign_key_exists?(:audit_logs, :items)
    add_index :audit_logs, :item_id unless index_exists?(:audit_logs, :item_id)
    
    # Now we can safely drop the unwanted tables
    drop_table :inventories if table_exists?(:inventories)
    drop_table :products if table_exists?(:products)
    drop_table :product_types if table_exists?(:product_types)
    
    # Modify users table
    if column_exists?(:users, :jwt_token)
      remove_column :users, :jwt_token
      remove_column :users, :ip_address
    end
    change_column :users, :role, :string, limit: 20
    add_index :users, :email, unique: true unless index_exists?(:users, :email, unique: true)
    
    # Modify items table - add weight column
    add_column :items, :weight, :float unless column_exists?(:items, :weight)
    
    # Modify audit_logs action column 
    change_column :audit_logs, :action, :string, limit: 50
    
    # Remove updated_at from usages and photos (not in your design)
    remove_column :usages, :updated_at if column_exists?(:usages, :updated_at)
    remove_column :photos, :updated_at if column_exists?(:photos, :updated_at)
  end
  
  def down
    # Recreate the dropped tables (basic structure)
    create_table :product_types do |t|
      t.string :name
      t.text :description
      t.timestamps
    end
    
    create_table :products do |t|
      t.string :name
      t.string :brand
      t.integer :quantity
      t.decimal :price
      t.timestamps
    end
    
    create_table :inventories do |t|
      t.references :product, null: false, foreign_key: true
      t.integer :quantity
      t.string :location
      t.string :batch_number
      t.timestamps
    end
    
    # Reverse other changes
    add_column :users, :jwt_token, :string
    add_column :users, :ip_address, :string
    remove_index :users, :email if index_exists?(:users, :email, unique: true)
    
    remove_column :items, :weight if column_exists?(:items, :weight)
    
    remove_foreign_key :audit_logs, :items if foreign_key_exists?(:audit_logs, :items)
    remove_column :audit_logs, :item_id if column_exists?(:audit_logs, :item_id)
    add_column :audit_logs, :product_id, :bigint
    add_foreign_key :audit_logs, :products
    
    add_column :usages, :updated_at, :datetime
    add_column :photos, :updated_at, :datetime
  end
end
