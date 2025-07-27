class CreateUsages < ActiveRecord::Migration[8.0]
  def change
    create_table :usages do |t|
      t.references :item, null: false, foreign_key: true
      t.integer :change_amount

      t.timestamps
    end
  end
end
