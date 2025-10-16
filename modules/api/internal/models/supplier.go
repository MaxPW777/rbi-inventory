package models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Name          string    `json:"name" db:"name" binding:"required"`
	ContactPerson *string   `json:"contact_person,omitempty" db:"contact_person"`
	Email         *string   `json:"email,omitempty" db:"email"`
	Phone         *string   `json:"phone,omitempty" db:"phone"`
	Address       *string   `json:"address,omitempty" db:"address"`
	Website       *string   `json:"website,omitempty" db:"website"`
	Notes         *string   `json:"notes,omitempty" db:"notes"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type SupplierProduct struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	SupplierID           uuid.UUID `json:"supplier_id" db:"supplier_id" binding:"required"`
	ProductID            uuid.UUID `json:"product_id" db:"product_id" binding:"required"`
	SupplierSKU          *string   `json:"supplier_sku,omitempty" db:"supplier_sku"`
	CostPerUnit          *float64  `json:"cost_per_unit,omitempty" db:"cost_per_unit"`
	LeadTimeDays         *int      `json:"lead_time_days,omitempty" db:"lead_time_days"`
	MinimumOrderQuantity *int      `json:"minimum_order_quantity,omitempty" db:"minimum_order_quantity"`
	IsPreferred          bool      `json:"is_preferred" db:"is_preferred"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}
