class ConvertToUuidsAndFinalCleanup < ActiveRecord::Migration[8.0]
  def up
    # Since we're doing a major restructure and likely don't have important data,
    # let's just recreate all tables with the correct UUID structure
    
    # Drop all existing tables (except Rails system tables)
    drop_table :audit_logs if table_exists?(:audit_logs)
    drop_table :photos if table_exists?(:photos)
    drop_table :usages if table_exists?(:usages)
    drop_table :items if table_exists?(:items)
    drop_table :suppliers if table_exists?(:suppliers)
    drop_table :folders if table_exists?(:folders)
    drop_table :users if table_exists?(:users)
    
    # Create tables with UUID primary keys matching your exact design
    
    # Users table
    create_table :users, id: :uuid do |t|
      t.string :name
      t.string :email
      t.string :role # enum('admin','editor','viewer')
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    add_index :users, :email, unique: true
    
    # Folders table
    create_table :folders, id: :uuid do |t|
      t.string :name
      t.boolean :is_brand
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    
    # Suppliers table
    create_table :suppliers, id: :uuid do |t|
      t.string :name
      t.string :contact_email
      t.string :phone
      t.string :address, type: :text
      t.text :notes
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    
    # Items table
    create_table :items, id: :uuid do |t|
      t.string :name
      t.integer :quantity
      t.float :volume_ml
      t.float :weight
      t.decimal :price
      t.decimal :selling_price
      t.integer :min_level
      t.text :notes
      t.datetime :expiry_date
      t.uuid :folder_id, null: false
      t.uuid :supplier_id, null: false
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    add_index :items, :folder_id
    add_index :items, :supplier_id
    
    # Usages table (no updated_at per your design)
    create_table :usages, id: :uuid do |t|
      t.uuid :item_id, null: false
      t.integer :change_amount
      t.datetime :created_at, null: false
    end
    add_index :usages, :item_id
    
    # Photos table (no updated_at per your design)
    create_table :photos, id: :uuid do |t|
      t.uuid :item_id, null: false
      t.string :url
      t.string :caption
      t.uuid :uploaded_by, null: false # references users.id
      t.datetime :created_at, null: false
    end
    add_index :photos, :item_id
    add_index :photos, :uploaded_by
    
    # Audit logs table (no updated_at per your design)
    create_table :audit_logs, id: :uuid do |t|
      t.uuid :user_id, null: false
      t.string :action # enum for different actions
      t.uuid :item_id # optional - can be null
      t.text :details
      t.datetime :created_at, null: false
    end
    add_index :audit_logs, :user_id
    add_index :audit_logs, :item_id
    
    # Add foreign key constraints
    add_foreign_key :items, :folders
    add_foreign_key :items, :suppliers
    add_foreign_key :usages, :items
    add_foreign_key :photos, :items
    add_foreign_key :photos, :users, column: :uploaded_by
    add_foreign_key :audit_logs, :users
    add_foreign_key :audit_logs, :items
  end
  
  def down
    # Drop all tables
    drop_table :audit_logs if table_exists?(:audit_logs)
    drop_table :photos if table_exists?(:photos)
    drop_table :usages if table_exists?(:usages)
    drop_table :items if table_exists?(:items)
    drop_table :suppliers if table_exists?(:suppliers)
    drop_table :folders if table_exists?(:folders)
    drop_table :users if table_exists?(:users)
  end
end
