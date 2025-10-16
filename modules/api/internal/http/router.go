package http

import (
	"net/http"

	"github.com/maximilianpw/rbi-inventory/internal/database"
	"github.com/maximilianpw/rbi-inventory/internal/handlers/users"

	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.Engine, db *database.DB) {
	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	v1 := r.Group("/api/v1")

	usersGroup := v1.Group("/users")
	users.BuildRoutes(usersGroup, db)
}
