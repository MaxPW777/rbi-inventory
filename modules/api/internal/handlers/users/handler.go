package users

import (
	"net/http"

	"github.com/maximilianpw/rbi-inventory/internal/database"
	"github.com/maximilianpw/rbi-inventory/internal/repository/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	repo users.UserRepository
}

func (handler *userHandler) handleGetUser(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return
	}

	user, err := handler.repo.GetByID(c, uuid)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) handleGetUsers(c *gin.Context) {
	users, err := handler.repo.List(c, []string{})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, users)
}

func NewHandler(repo *users.UserRepository) *userHandler {
	return &userHandler{repo: *repo}
}

func BuildRoutes(rg *gin.RouterGroup, db *database.DB) {
	userRepo := users.NewRepository(nil)
	handler := NewHandler(&userRepo)
	rg.GET("", handler.handleGetUsers)
	rg.GET("/:id", handler.handleGetUser)
}
