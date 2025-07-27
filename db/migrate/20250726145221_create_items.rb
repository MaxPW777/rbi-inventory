class CreateItems < ActiveRecord::Migration[8.0]
  def change
    create_table :items do |t|
      t.string :name
      t.integer :quantity
      t.float :volume_ml
      t.float :weight
      t.decimal :price
      t.decimal :selling_price
      t.integer :min_level
      t.text :notes
      t.datetime :expiry_date
      t.references :folder, null: false, foreign_key: true
      t.references :supplier, null: false, foreign_key: true

      t.timestamps
    end
  end
end
