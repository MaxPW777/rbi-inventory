package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Role      string     `db:"role"`
	IsActive  bool       `db:"is_active"`
	LastLogin *time.Time `db:"last_login"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

