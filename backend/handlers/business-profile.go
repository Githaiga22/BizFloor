package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"group7/models"

	"gorm.io/gorm"
)

type ProfileHandler struct {
	DB *gorm.DB
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var business models.Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

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
