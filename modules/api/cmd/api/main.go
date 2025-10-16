package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/maximilianpw/rbi-inventory/internal/config"
	"github.com/maximilianpw/rbi-inventory/internal/database"
	"github.com/maximilianpw/rbi-inventory/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	defer db.Close()

	log.Println("Database Connected successfully")

	r := gin.Default()
	http.BuildRouter(r, db)

	log.Printf("listening on port: %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
