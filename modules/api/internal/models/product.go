package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `db:"id"`
	Name        string     `db:"name"`
	ParentID    *uuid.UUID `db:"parent_id"`
	Description *string    `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

type Brand struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	Website     *string   `db:"website"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Product struct {
	ID                uuid.UUID  `db:"id"`
	SKU               string     `db:"sku"`
	Name              string     `db:"name"`
	Description       *string    `db:"description"`
	CategoryID        uuid.UUID  `db:"category_id"`
	BrandID           *uuid.UUID `db:"brand_id"`
	VolumeML          *int       `db:"volume_ml"`
	WeightKG          *float64   `db:"weight_kg"`
	DimensionsCM      *string    `db:"dimensions_cm"`
	StandardCost      *float64   `db:"standard_cost"`
	StandardPrice     *float64   `db:"standard_price"`
	MarkupPercentage  *float64   `db:"markup_percentage"`
	ReorderPoint      int        `db:"reorder_point"`
	PrimarySupplierID *uuid.UUID `db:"primary_supplier_id"`
	SupplierSKU       *string    `db:"supplier_sku"`
	IsActive          bool       `db:"is_active"`
	IsPerishable      bool       `db:"is_perishable"`
	Notes             *string    `db:"notes"`
	CreatedAt         time.Time  `db:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at"`
}

type Photo struct {
	ID         uuid.UUID  `db:"id"`
	ProductID  uuid.UUID  `db:"product_id"`
	URL        string     `db:"url"`
	Caption    *string    `db:"caption"`
	UploadedBy *uuid.UUID `db:"uploaded_by"`
	CreatedAt  time.Time  `db:"created_at"`
}
