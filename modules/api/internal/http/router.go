package http

import (
	"net/http"

	"rbi/api/internal/handlers/users"

	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.Engine) {
	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	v1 := r.Group("/api/v1")

	usersGroup := v1.Group("/users")
	users.BuildRoutes(usersGroup)
}
