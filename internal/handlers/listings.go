package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dev-sandip/hamrobazar-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListingHandler struct {
	DB *gorm.DB
}

func NewListingHandler(db *gorm.DB) *ListingHandler {
	return &ListingHandler{DB: db}
}

func (h *ListingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var listing models.Listing

	if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Don't allow client to control these fields.
	listing.ID = uuid.Nil
	listing.CreatedAt = time.Time{}

	if err := h.DB.Create(&listing).Error; err != nil {
		http.Error(w, "failed to create listing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(listing)
}

func (h *ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	var listings []models.Listing

	if err := h.DB.Order("created_at DESC").Find(&listings).Error; err != nil {
		http.Error(w, "failed to fetch listings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(listings)
}
