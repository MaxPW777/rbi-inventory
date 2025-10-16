package models

import (
	"time"

	"github.com/google/uuid"
)

type StockMovementReason string

const (
	StockMovementPurchaseReceive  StockMovementReason = "PURCHASE_RECEIVE"
	StockMovementSale             StockMovementReason = "SALE"
	StockMovementWaste            StockMovementReason = "WASTE"
	StockMovementDamaged          StockMovementReason = "DAMAGED"
	StockMovementExpired          StockMovementReason = "EXPIRED"
	StockMovementCountCorrection  StockMovementReason = "COUNT_CORRECTION"
	StockMovementReturnFromClient StockMovementReason = "RETURN_FROM_CLIENT"
	StockMovementReturnToSupplier StockMovementReason = "RETURN_TO_SUPPLIER"
	StockMovementInternalTransfer StockMovementReason = "INTERNAL_TRANSFER"
)

type LocationType string

const (
	LocationTypeWarehouse LocationType = "WAREHOUSE"
	LocationTypeSupplier  LocationType = "SUPPLIER"
	LocationTypeInTransit LocationType = "IN_TRANSIT"
	LocationTypeClient    LocationType = "CLIENT"
)

type Location struct {
	ID            uuid.UUID    `json:"id" db:"id"`
	Name          string       `json:"name" db:"name" binding:"required"`
	Type          LocationType `json:"type" db:"type" binding:"required"`
	Address       *string      `json:"address,omitempty" db:"address"`
	ContactPerson *string      `json:"contact_person,omitempty" db:"contact_person"`
	Phone         *string      `json:"phone,omitempty" db:"phone"`
	IsActive      bool         `json:"is_active" db:"is_active"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
}

type Inventory struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	ProductID    uuid.UUID  `json:"product_id" db:"product_id" binding:"required"`
	LocationID   uuid.UUID  `json:"location_id" db:"location_id" binding:"required"`
	Quantity     int        `json:"quantity" db:"quantity"`
	BatchNumber  *string    `json:"batch_number,omitempty" db:"batch_number"`
	ExpiryDate   *time.Time `json:"expiry_date,omitempty" db:"expiry_date"`
	CostPerUnit  *float64   `json:"cost_per_unit,omitempty" db:"cost_per_unit"`
	ReceivedDate *time.Time `json:"received_date,omitempty" db:"received_date"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

type StockMovement struct {
	ID              uuid.UUID           `json:"id" db:"id"`
	ProductID       uuid.UUID           `json:"product_id" db:"product_id" binding:"required"`
	FromLocationID  *uuid.UUID          `json:"from_location_id,omitempty" db:"from_location_id"`
	ToLocationID    *uuid.UUID          `json:"to_location_id,omitempty" db:"to_location_id"`
	Quantity        int                 `json:"quantity" db:"quantity" binding:"required"`
	Reason          StockMovementReason `json:"reason" db:"reason" binding:"required"`
	OrderID         *uuid.UUID          `json:"order_id,omitempty" db:"order_id"`
	ReferenceNumber *string             `json:"reference_number,omitempty" db:"reference_number"`
	CostPerUnit     *float64            `json:"cost_per_unit,omitempty" db:"cost_per_unit"`
	UserID          uuid.UUID           `json:"user_id" db:"user_id" binding:"required"`
	Notes           *string             `json:"notes,omitempty" db:"notes"`
	CreatedAt       time.Time           `json:"created_at" db:"created_at"`
}
