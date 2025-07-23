class Product < ApplicationRecord
  validates :name, presence: true
  belongs_to :product_type, optional: true
  has_many :inventories
  has_many :audit_logs
end
