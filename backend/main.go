package main

import (
	"log"
	"net/http"

	"group7/handlers"

	"group7/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically migrate models (optional)
	if err := db.AutoMigrate(
		&models.Business{},
		&models.Service{},
		&models.Booking{},
		&models.Payment{},
	); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Initialize the ServeMux
	mux := http.NewServeMux()

	// Register handlers
	handlers.RegisterHandlers(mux, db)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
