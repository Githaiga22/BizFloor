package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"group7/auth"
	"group7/models"

	"gorm.io/gorm"
)

type BusinessProfileHandler struct {
	DB *gorm.DB
}

func (h *BusinessProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
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

	var business models.Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set OwnerID to the logged-in user's ID
	business.OwnerID = userID

	// Set timestamps for creation
	business.CreatedAt = time.Now()
	business.UpdatedAt = time.Now()

	if err := h.DB.Create(&business).Error; err != nil {
		http.Error(w, "Failed to create business profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(business); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
