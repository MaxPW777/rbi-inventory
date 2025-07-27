class CreateSuppliers < ActiveRecord::Migration[8.0]
  def change
    create_table :suppliers do |t|
      t.string :name
      t.string :contact_email
      t.string :phone
      t.string :address
      t.text :notes

      t.timestamps
    end
  end
end
