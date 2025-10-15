package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maximilianpw/rbi-inventory/internal/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	http.BuildRouter(r)

	log.Printf("listening on port: %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
