package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"rbi/api/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	http.BuildRouter(r)

	log.Printf("listening on port: %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
