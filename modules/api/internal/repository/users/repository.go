package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/maximilianpw/rbi-inventory/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	List(ctx context.Context, filters []string) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *sqlx.DB
}

var users = []models.User{
	{ID: uuid.Must(uuid.NewRandom()), Name: "Max"},
	{ID: uuid.Must(uuid.NewRandom()), Name: "Lukas"},
}

// Create implements Repository.
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetByID implements Repository.
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return &users[0], nil
}

// List implements Repository.
func (r *userRepository) List(ctx context.Context, filters []string) ([]*models.User, error) {
	result := make([]*models.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	return result, nil
}

// Update implements Repository.
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

func NewRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
