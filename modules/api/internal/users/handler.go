package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	name string `json:"name"`
}

var users = []User{
	{ID: "1", name: "max"},
	{ID: "2", name: "lukas"},
}

func handleGetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func handleGetUser(c *gin.Context) {
	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}

func BuildRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleGetUsers)
	rg.GET(":id", handleGetUser)
}
