class Item < ApplicationRecord
  belongs_to :folder
  belongs_to :supplier
end
