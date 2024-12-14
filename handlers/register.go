package handlers

import (
	"net/http"

	"group7/auth"

	"gorm.io/gorm"
)

func RegisterHandlers(mux *http.ServeMux, db *gorm.DB) {
	authHandler := auth.NewAuthHandler(db)
	profileHandler := &ProfileHandler{DB: db}
	businessProfileHandler := &BusinessProfileHandler{DB: db}
	serviceHandler := &ServiceHandler{DB: db}

	// Auth routes
	mux.HandleFunc("/api/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)

	// Initialize booking handler
	bookingHandler := NewBookingHandler(db)

	// Add booking routes
	mux.Handle("/api/bookings", authHandler.AuthMiddleware(http.HandlerFunc(bookingHandler.CreateBooking)))
	mux.Handle("/api/bookings/list", authHandler.AuthMiddleware(http.HandlerFunc(bookingHandler.GetBookings)))
	mux.Handle("/api/bookings/status", authHandler.AuthMiddleware(http.HandlerFunc(bookingHandler.UpdateBookingStatus)))
	// Protected routes
	mux.Handle("/api/profile", authHandler.AuthMiddleware(http.HandlerFunc(profileHandler.profileHandler)))
	mux.Handle("/api/create-profile", authHandler.AuthMiddleware(http.HandlerFunc(businessProfileHandler.CreateProfile)))
	mux.Handle("/api/add-service", authHandler.AuthMiddleware(http.HandlerFunc(serviceHandler.AddServiceHandler)))


	// Public page routes
	mux.HandleFunc("/", serveTemplate("index.html"))
	mux.HandleFunc("/login", serveTemplate("login.html"))
	mux.HandleFunc("/signup", serveTemplate("signup.html"))
	mux.Handle("/business-dashboard", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("business-dashboard.html"))))
	mux.Handle("/customer-dashboard", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("customer-dashboard.html"))))
	mux.Handle("/create-profile", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("create-business-profile.html"))))
	mux.Handle("/add-service", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("add-service.html"))))
	mux.Handle("/withdraw-funds", authHandler.AuthMiddleware(http.HandlerFunc(serveTemplate("withdraw-funds.html"))))
}

func serveTemplate(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/"+templateName)
	}
}
