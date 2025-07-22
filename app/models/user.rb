class User < ApplicationRecord
  has_many :audit_logs
end
