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
	ID            uuid.UUID    `db:"id"`
	Name          string       `db:"name"`
	Type          LocationType `db:"type"`
	Address       *string      `db:"address"`
	ContactPerson *string      `db:"contact_person"`
	Phone         *string      `db:"phone"`
	IsActive      bool         `db:"is_active"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
}

type Inventory struct {
	ID           uuid.UUID  `db:"id"`
	ProductID    uuid.UUID  `db:"product_id"`
	LocationID   uuid.UUID  `db:"location_id"`
	Quantity     int        `db:"quantity"`
	BatchNumber  *string    `db:"batch_number"`
	ExpiryDate   *time.Time `db:"expiry_date"`
	CostPerUnit  *float64   `db:"cost_per_unit"`
	ReceivedDate *time.Time `db:"received_date"`
	UpdatedAt    time.Time  `db:"updated_at"`
}

type StockMovement struct {
	ID              uuid.UUID           `db:"id"`
	ProductID       uuid.UUID           `db:"product_id"`
	FromLocationID  *uuid.UUID          `db:"from_location_id"`
	ToLocationID    *uuid.UUID          `db:"to_location_id"`
	Quantity        int                 `db:"quantity"`
	Reason          StockMovementReason `db:"reason"`
	OrderID         *uuid.UUID          `db:"order_id"`
	ReferenceNumber *string             `db:"reference_number"`
	CostPerUnit     *float64            `db:"cost_per_unit"`
	UserID          uuid.UUID           `db:"user_id"`
	Notes           *string             `db:"notes"`
	CreatedAt       time.Time           `db:"created_at"`
}
