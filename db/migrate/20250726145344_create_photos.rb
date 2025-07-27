class CreatePhotos < ActiveRecord::Migration[8.0]
  def change
    create_table :photos do |t|
      t.references :item, null: false, foreign_key: true
      t.string :url
      t.string :caption
      t.references :uploaded_by, null: false, foreign_key: { to_table: :users }

      t.timestamps
    end
  end
end
