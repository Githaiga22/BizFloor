package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"group7/auth"
	"group7/db"
	"group7/models"
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

	// Initialize auth handler
	authHandler := auth.NewAuthHandler(database)

	// Setup routes
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Auth API routes
	mux.HandleFunc("/api/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)

	// Protected API routes
	mux.Handle("/api/profile", authHandler.AuthMiddleware(http.HandlerFunc(profileHandler)))

	// Page routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	})
	mux.HandleFunc("/login", serveTemplate("login.html"))
	mux.HandleFunc("/signup", serveTemplate("signup.html"))
	mux.Handle("/business-dashboard", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("business-dashboard.html"))))
	mux.Handle("/customer-dashboard", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("customer-dashboard.html"))))

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func serveTemplate(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, templateName, nil)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	session := r.Context().Value("session").(*auth.Session)

	if session == nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Fetch user from database to get is_business_owner status
	var user models.User
	if err := database.Where("id = ?", userID).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := auth.User{
		ID:              userID,
		Name:            session.UserName,
		Email:           session.UserEmail,
		IsBusinessOwner: user.IsBusinessOwner,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
