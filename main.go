package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"group7/db"
	"group7/handlers"

	"gorm.io/gorm"
)

var (
	tmpl     *template.Template
	database *gorm.DB
)

func init() {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Parse templates using current working directory
	tmpl, err = template.ParseGlob(filepath.Join(currentDir, "templates", "*.html"))
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
	}
}

func main() {
	// Initialize database
	var err error
	database, err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup routes
	mux := http.NewServeMux()

	// Register all handlers
	handlers.RegisterHandlers(mux, database)

	// Start the server
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
