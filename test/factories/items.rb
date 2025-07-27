FactoryBot.define do
  factory :item do
    name { "MyString" }
    quantity { 1 }
    volume_ml { 1.5 }
    weight { 1.5 }
    price { "9.99" }
    selling_price { "9.99" }
    min_level { 1 }
    notes { "MyText" }
    expiry_date { "2025-07-26 14:52:21" }
    folder { nil }
    supplier { nil }
  end
end
