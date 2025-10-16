package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditAction string

const (
	AuditActionCreate         AuditAction = "CREATE"
	AuditActionUpdate         AuditAction = "UPDATE"
	AuditActionDelete         AuditAction = "DELETE"
	AuditActionAdjustQuantity AuditAction = "ADJUST_QUANTITY"
	AuditActionAddPhoto       AuditAction = "ADD_PHOTO"
	AuditActionStatusChange   AuditAction = "STATUS_CHANGE"
)

type AuditLog struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	UserID     *uuid.UUID  `json:"user_id,omitempty" db:"user_id"`
	Action     AuditAction `json:"action" db:"action" binding:"required"`
	EntityType string      `json:"entity_type" db:"entity_type" binding:"required"`
	EntityID   uuid.UUID   `json:"entity_id" db:"entity_id" binding:"required"`
	Changes    *string     `json:"changes,omitempty" db:"changes"` // JSONB stored as string
	IPAddress  *string     `json:"ip_address,omitempty" db:"ip_address"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
}
