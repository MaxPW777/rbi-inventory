package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Request.Response.Status
		_ = latency
		_ = status
		_ = method
		_ = path
	}
}
