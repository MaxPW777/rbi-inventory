package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"scooby": "doo"})
}

func BuildRoutes(rg *gin.RouterGroup) {
	rg.GET("/users", handleGetUsers)
}
