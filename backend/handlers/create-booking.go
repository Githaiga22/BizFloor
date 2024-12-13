// File: handlers/create_booking.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"group7/models"

	"gorm.io/gorm"
)

type BookingHandler struct {
	DB *gorm.DB
}

func (h *BookingHandler) CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the ServiceID
	var service models.Service
	if err := h.DB.First(&service, booking.ServiceID).Error; err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Validate the ClientID
	var client models.User
	if err := h.DB.First(&client, booking.ClientID).Error; err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// Validate the BusinessID
	var business models.Business
	if err := h.DB.First(&business, booking.BusinessID).Error; err != nil {
		http.Error(w, "Business not found", http.StatusNotFound)
		return
	}

	// Set timestamps for creation
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	// Set default status if not provided
	if booking.Status == "" {
		booking.Status = "pending"
	}

	// Create the booking
	if err := h.DB.Create(&booking).Error; err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	// Return the booking as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(booking); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
