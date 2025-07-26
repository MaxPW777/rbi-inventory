class Photo < ApplicationRecord
  belongs_to :item
  belongs_to :uploaded_by
end
