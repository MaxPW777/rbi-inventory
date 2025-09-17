package main

import (
	"log"

	"github.com/gin-gonic/gin"

	httpadapter "rbi/api/internal/adapter/http"
	"rbi/api/internal/adapter/memory"
	application "rbi/api/internal/core/app"
)

func main() {
	repository := memory.NewGreetingRepository()
	service := application.NewGreetingService(repository)
	handler := httpadapter.NewGreetingHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}
