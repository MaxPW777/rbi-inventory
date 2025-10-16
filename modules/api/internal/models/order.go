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
	ID                  uuid.UUID   `db:"id"`
	OrderNumber         string      `db:"order_number"`
	ClientID            uuid.UUID   `db:"client_id"`
	Status              OrderStatus `db:"status"`
	DeliveryDeadline    *time.Time  `db:"delivery_deadline"`
	DeliveryAddress     string      `db:"delivery_address"`
	YachtName           *string     `db:"yacht_name"`
	SpecialInstructions *string     `db:"special_instructions"`
	TotalAmount         float64     `db:"total_amount"`
	CreatedBy           uuid.UUID   `db:"created_by"`
	ConfirmedAt         *time.Time  `db:"confirmed_at"`
	ShippedAt           *time.Time  `db:"shipped_at"`
	DeliveredAt         *time.Time  `db:"delivered_at"`
	CreatedAt           time.Time   `db:"created_at"`
	UpdatedAt           time.Time   `db:"updated_at"`
}

type OrderItem struct {
	ID             uuid.UUID `db:"id"`
	OrderID        uuid.UUID `db:"order_id"`
	ProductID      uuid.UUID `db:"product_id"`
	Quantity       int       `db:"quantity"`
	UnitPrice      float64   `db:"unit_price"`
	Subtotal       float64   `db:"subtotal"`
	Notes          *string   `db:"notes"`
	QuantityPicked int       `db:"quantity_picked"`
	QuantityPacked int       `db:"quantity_packed"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
