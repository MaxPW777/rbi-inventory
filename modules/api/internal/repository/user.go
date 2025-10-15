package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"rbi/api/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, client *models.User) error
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	List(ctx context.Context, filters []string) ([]*models.User, error)
	Update(ctx context.Context, client *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepo struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

var users = []models.User{
	{ID: uuid.Must(uuid.NewRandom()), Name: "Max"},
	{ID: uuid.Must(uuid.NewRandom()), Name: "Lukas"},
}

func (repo *userRepo) List() []models.User {
	return users
}

func HandleGetUser(id uuid.UUID) models.User {
	return users[0]
}
