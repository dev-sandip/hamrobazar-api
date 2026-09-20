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

// Create godoc
//
//	@Summary		Create a listing
//	@Description	Create a new property listing
//	@Tags		Listings
//	@Accept		json
//	@Produce	json
//	@Param		listing	body	CreateListingRequest	true	"Listing"
//	@Success	201	{object}	models.Listing
//	@Failure	400	{string}	string
//	@Failure	500	{string}	string
//	@Router		/listings [post]

func (lh *ListingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var listing models.Listing

	if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Don't allow client to control these fields.
	listing.ID = uuid.Nil
	listing.CreatedAt = time.Time{}

	if err := lh.DB.Create(&listing).Error; err != nil {
		http.Error(w, "failed to create listing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(listing)
}

// List godoc
//
//	@Summary		List all listings
//	@Description	List all property listings
//	@Tags		Listings
//	@Produce	json
//	@Success	200	{array}	models.Listing
//	@Failure	500	{string}	string
//	@Router		/listings [get]
func (lh *ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	var listings []models.Listing

	if err := lh.DB.Order("created_at DESC").Find(&listings).Error; err != nil {
		http.Error(w, "failed to fetch listings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(listings)
}

// Delete godoc
//
//	@Summary		Delete a listing
//	@Description	Delete a listing by UUID
//	@Tags		Listings
//	@Produce	plain
//	@Param		id	path	string	true	"Listing UUID"
//	@Success	204
//	@Failure	400	{string}	string
//	@Failure	404	{string}	string
//	@Failure	500	{string}	string
//	@Router		/listings/{id} [delete]

func (lh *ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	listingId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid listing id", http.StatusBadRequest)
		return
	}

	result := lh.DB.Delete(&models.Listing{ID: listingId})
	if result.Error != nil {
		http.Error(w, "Failed to delete listing", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
