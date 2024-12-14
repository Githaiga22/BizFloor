// File: handlers/service_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"group7/auth"
	"group7/models"

	"gorm.io/gorm"
)

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) AddServiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session := r.Context().Value("session").(*auth.Session)

	if session == nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Extract userID from context (set by the AuthMiddleware)
	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Check if the user is a business owner
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Ensure the user is a business owner
	if !user.IsBusinessOwner {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Query the Business table to get the business owned by this user
	var business models.Business
	if err := h.DB.Where("owner_id = ?", userID).First(&business).Error; err != nil {
		http.Error(w, "Business not found for this user", http.StatusNotFound)
		return
	}

	// Now that we have the business, set the BusinessID in the service
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the BusinessID to the ID of the business found
	service.BusinessID = business.ID

	// Set timestamps for creation
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	if err := h.DB.Create(&service).Error; err != nil {
		http.Error(w, "Failed to create service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(service); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
