package httpadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rbi/api/internal/port/inbound"
)

// GreetingHandler wires the inbound port to an HTTP transport.
type GreetingHandler struct {
	service inbound.GreetingUseCase
}

func NewGreetingHandler(service inbound.GreetingUseCase) *GreetingHandler {
	return &GreetingHandler{service: service}
}

func (h *GreetingHandler) RegisterRoutes(router gin.IRouter) {
	router.GET("/greet/:name", h.handleGreet)
}

func (h *GreetingHandler) handleGreet(c *gin.Context) {
	name := c.Param("name")

	greeting, err := h.service.Greet(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    greeting.Name,
		"message": greeting.Message,
	})
}
