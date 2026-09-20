package models

import (
	"time"

	"github.com/google/uuid"
)

type Listing struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Price       int64     `gorm:"not null"`
	City        string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
}
