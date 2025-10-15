package users

import (
	"net/http"

	"github.com/maximilianpw/rbi-inventory/internal/repository/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func handleGetUser(c *gin.Context) {
	userRepo := users.NewRepository(nil)
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return
	}

	user, err := userRepo.GetByID(c, uuid)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func handleGetUsers(c *gin.Context) {
	userRepo := users.NewRepository(nil)
	users, err := userRepo.List(c, []string{})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, users)
}

func BuildRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleGetUsers)
	rg.GET("/:id", handleGetUser)
}
