package users

import (
	"context"
	"net/http"
	"time"

	"github.com/maximilianpw/rbi-inventory/internal/database"
	"github.com/maximilianpw/rbi-inventory/internal/handlers/users/dtos"
	"github.com/maximilianpw/rbi-inventory/internal/repository/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	repo users.UserRepository
}

func (handler *userHandler) HandleGetUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := handler.repo.GetByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) HandleGetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	users, err := handler.repo.List(ctx, []string{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	responses := make([]dtos.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, dtos.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			IsActive:  user.IsActive,
			LastLogin: user.LastLogin,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func NewHandler(db *database.DB) *userHandler {
	userRepo := users.NewRepository(db.DB)
	return &userHandler{repo: userRepo}
}

func BuildRoutes(rg *gin.RouterGroup, db *database.DB) {
	handler := NewHandler(db)
	rg.GET("", handler.HandleGetUsers)
	rg.GET("/:id", handler.HandleGetUser)
}
