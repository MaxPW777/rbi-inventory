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
	ID                     uuid.UUID    `db:"id"`
	CompanyName            string       `db:"company_name"`
	YachtName              *string      `db:"yacht_name"`
	ContactPerson          string       `db:"contact_person"`
	Email                  string       `db:"email"`
	Phone                  *string      `db:"phone"`
	BillingAddress         *string      `db:"billing_address"`
	DefaultDeliveryAddress *string      `db:"default_delivery_address"`
	AccountStatus          ClientStatus `db:"account_status"`
	PaymentTerms           *string      `db:"payment_terms"`
	CreditLimit            *float64     `db:"credit_limit"`
	Notes                  *string      `db:"notes"`
	CreatedAt              time.Time    `db:"created_at"`
	UpdatedAt              time.Time    `db:"updated_at"`
}
