class CreateAuditLogs < ActiveRecord::Migration[8.0]
  def change
    create_table :audit_logs do |t|
      t.references :user, null: false, foreign_key: true
      t.integer :action
      t.references :item, null: false, foreign_key: true
      t.text :details

      t.timestamps
    end
  end
end
