package models

import (
	"time"

	"github.com/google/uuid"
)

type ClientStatus string

const (
	ClientStatusActive    ClientStatus = "ACTIVE"
	ClientStatusSuspended ClientStatus = "SUSPENDED"
	ClientStatusInactive  ClientStatus = "INACTIVE"
)

type Client struct {
	ID                     uuid.UUID    `json:"id" db:"id"`
	CompanyName            string       `json:"company_name" db:"company_name" binding:"required"`
	YachtName              *string      `json:"yacht_name,omitempty" db:"yacht_name"`
	ContactPerson          string       `json:"contact_person" db:"contact_person" binding:"required"`
	Email                  string       `json:"email" db:"email" binding:"required,email"`
	Phone                  *string      `json:"phone,omitempty" db:"phone"`
	BillingAddress         *string      `json:"billing_address,omitempty" db:"billing_address"`
	DefaultDeliveryAddress *string      `json:"default_delivery_address,omitempty" db:"default_delivery_address"`
	AccountStatus          ClientStatus `json:"account_status" db:"account_status"`
	PaymentTerms           *string      `json:"payment_terms,omitempty" db:"payment_terms"`
	CreditLimit            *float64     `json:"credit_limit,omitempty" db:"credit_limit"`
	Notes                  *string      `json:"notes,omitempty" db:"notes"`
	CreatedAt              time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at" db:"updated_at"`
}
