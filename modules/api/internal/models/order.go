package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusDraft     OrderStatus = "DRAFT"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusSourcing  OrderStatus = "SOURCING"
	OrderStatusPicking   OrderStatus = "PICKING"
	OrderStatusPacked    OrderStatus = "PACKED"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusOnHold    OrderStatus = "ON_HOLD"
)

type Order struct {
	ID                  uuid.UUID   `json:"id" db:"id"`
	OrderNumber         string      `json:"order_number" db:"order_number" binding:"required"`
	ClientID            uuid.UUID   `json:"client_id" db:"client_id" binding:"required"`
	Status              OrderStatus `json:"status" db:"status"`
	DeliveryDeadline    *time.Time  `json:"delivery_deadline,omitempty" db:"delivery_deadline"`
	DeliveryAddress     string      `json:"delivery_address" db:"delivery_address" binding:"required"`
	YachtName           *string     `json:"yacht_name,omitempty" db:"yacht_name"`
	SpecialInstructions *string     `json:"special_instructions,omitempty" db:"special_instructions"`
	TotalAmount         float64     `json:"total_amount" db:"total_amount"`
	CreatedBy           uuid.UUID   `json:"created_by" db:"created_by" binding:"required"`
	ConfirmedAt         *time.Time  `json:"confirmed_at,omitempty" db:"confirmed_at"`
	ShippedAt           *time.Time  `json:"shipped_at,omitempty" db:"shipped_at"`
	DeliveredAt         *time.Time  `json:"delivered_at,omitempty" db:"delivered_at"`
	CreatedAt           time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at" db:"updated_at"`
}

type OrderItem struct {
	ID             uuid.UUID `json:"id" db:"id"`
	OrderID        uuid.UUID `json:"order_id" db:"order_id" binding:"required"`
	ProductID      uuid.UUID `json:"product_id" db:"product_id" binding:"required"`
	Quantity       int       `json:"quantity" db:"quantity" binding:"required,min=1"`
	UnitPrice      float64   `json:"unit_price" db:"unit_price" binding:"required,min=0"`
	Subtotal       float64   `json:"subtotal" db:"subtotal"`
	Notes          *string   `json:"notes,omitempty" db:"notes"`
	QuantityPicked int       `json:"quantity_picked" db:"quantity_picked"`
	QuantityPacked int       `json:"quantity_packed" db:"quantity_packed"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
