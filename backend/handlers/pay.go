// handlers/pay.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"group7/models"

	"gorm.io/gorm"
)

type PaymentHandler struct {
	DB *gorm.DB
}

func (h *PaymentHandler) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the BookingID
	var booking models.Booking
	if err := h.DB.First(&booking, payment.BookingID).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Set the PaymentTime if not provided
	if payment.PaymentTime.IsZero() {
		payment.PaymentTime = time.Now()
	}

	// Set default status if not provided
	if payment.Status == "" {
		payment.Status = "pending"
	}

	// Create the payment record
	if err := h.DB.Create(&payment).Error; err != nil {
		http.Error(w, "Failed to create payment", http.StatusInternalServerError)
		return
	}

	// Return the payment object as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(payment); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
