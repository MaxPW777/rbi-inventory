package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name" binding:"required"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
	Description *string    `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type Brand struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description *string   `json:"description,omitempty" db:"description"`
	Website     *string   `json:"website,omitempty" db:"website"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Product struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	SKU               string     `json:"sku" db:"sku" binding:"required"`
	Name              string     `json:"name" db:"name" binding:"required"`
	Description       *string    `json:"description,omitempty" db:"description"`
	CategoryID        uuid.UUID  `json:"category_id" db:"category_id" binding:"required"`
	BrandID           *uuid.UUID `json:"brand_id,omitempty" db:"brand_id"`
	VolumeML          *int       `json:"volume_ml,omitempty" db:"volume_ml"`
	WeightKG          *float64   `json:"weight_kg,omitempty" db:"weight_kg"`
	DimensionsCM      *string    `json:"dimensions_cm,omitempty" db:"dimensions_cm"`
	StandardCost      *float64   `json:"standard_cost,omitempty" db:"standard_cost"`
	StandardPrice     *float64   `json:"standard_price,omitempty" db:"standard_price"`
	MarkupPercentage  *float64   `json:"markup_percentage,omitempty" db:"markup_percentage"`
	ReorderPoint      int        `json:"reorder_point" db:"reorder_point"`
	PrimarySupplierID *uuid.UUID `json:"primary_supplier_id,omitempty" db:"primary_supplier_id"`
	SupplierSKU       *string    `json:"supplier_sku,omitempty" db:"supplier_sku"`
	IsActive          bool       `json:"is_active" db:"is_active"`
	IsPerishable      bool       `json:"is_perishable" db:"is_perishable"`
	Notes             *string    `json:"notes,omitempty" db:"notes"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

type Photo struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	ProductID  uuid.UUID  `json:"product_id" db:"product_id" binding:"required"`
	URL        string     `json:"url" db:"url" binding:"required"`
	Caption    *string    `json:"caption,omitempty" db:"caption"`
	UploadedBy *uuid.UUID `json:"uploaded_by,omitempty" db:"uploaded_by"`
	CreatedAttime.Time  `json:"created_at" db:"created_at"`
}
