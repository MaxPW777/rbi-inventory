package users

import (
	"net/http"

	"rbi/api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func handleGetUser(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return
	}

	user := repository.HandleGetUser(uuid)
	c.JSON(http.StatusFound, user)
}

func handleGetUsers(c *gin.Context) {
	users := repository.HandleGetUsers()

	c.JSON(http.StatusFound, users)
}

func BuildRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleGetUsers)
	rg.GET(":id", handleGetUser)
}
