package db

import (
	"fmt"
	"time"

	"github.com/dev-sandip/hamrobazar-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	if err := db.AutoMigrate(&models.Listing{}); err != nil {
		return nil, fmt.Errorf("db.AutoMigrate: %w", err)
	}

	return db, nil
}
