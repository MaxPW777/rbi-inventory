package users

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"rbi/api/internal/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	List(ctx context.Context, filters []string) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *sql.DB
}

var users = []models.User{
	{ID: uuid.Must(uuid.NewRandom()), Name: "Max"},
	{ID: uuid.Must(uuid.NewRandom()), Name: "Lukas"},
}

// Create implements Repository.
func (r *repository) Create(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetByID implements Repository.
func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return &users[0], nil
}

// List implements Repository.
func (r *repository) List(ctx context.Context, filters []string) ([]*models.User, error) {
	result := make([]*models.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	return result, nil
}

// Update implements Repository.
func (r *repository) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
