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
	ID         uuid.UUID   `db:"id"`
	UserID     *uuid.UUID  `db:"user_id"`
	Action     AuditAction `db:"action"`
	EntityType string      `db:"entity_type"`
	EntityID   uuid.UUID   `db:"entity_id"`
	Changes    *string     `db:"changes"` // JSONB stored as string
	IPAddress  *string     `db:"ip_address"`
	CreatedAt  time.Time   `db:"created_at"`
}
