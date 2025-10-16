package models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	ContactPerson *string   `db:"contact_person"`
	Email         *string   `db:"email"`
	Phone         *string   `db:"phone"`
	Address       *string   `db:"address"`
	Website       *string   `db:"website"`
	Notes         *string   `db:"notes"`
	IsActive      bool      `db:"is_active"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type SupplierProduct struct {
	ID                   uuid.UUID `db:"id"`
	SupplierID           uuid.UUID `db:"supplier_id"`
	ProductID            uuid.UUID `db:"product_id"`
	SupplierSKU          *string   `db:"supplier_sku"`
	CostPerUnit          *float64  `db:"cost_per_unit"`
	LeadTimeDays         *int      `db:"lead_time_days"`
	MinimumOrderQuantity *int      `db:"minimum_order_quantity"`
	IsPreferred          bool      `db:"is_preferred"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
