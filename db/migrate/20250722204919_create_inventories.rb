class CreateInventories < ActiveRecord::Migration[7.2]
  def change
    create_table :inventories do |t|
      t.references :product, null: false, foreign_key: true
      t.integer :quantity
      t.string :location
      t.string :batch_number

      t.timestamps
    end
  end
end
