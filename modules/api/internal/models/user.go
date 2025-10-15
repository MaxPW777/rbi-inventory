package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      string     `json:"name" db:"name" binding:"required"`
	Email     string     `json:"email" db:"email" binding:"required,email"`
	Role      string     `json:"role" db:"role" binding:"required"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty" db:"last_login"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

