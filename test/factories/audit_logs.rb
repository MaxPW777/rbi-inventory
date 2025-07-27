FactoryBot.define do
  factory :audit_log do
    user { nil }
    action { 1 }
    item { nil }
    details { "MyText" }
  end
end
