package app

import (
	"time"

	"gorm.io/gorm"
)

// Repository represents the repository layer
type Repository struct {
	DB *gorm.DB
}

// InitRepository initializes the repository component
func InitRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

// CreateMasterUrl creates a new master URL record (short url key, target url, created date) in database
func (repository *Repository) CreateMasterUrl(key string, targetUrl string) (*MasterUrl, error) {
	// Create new record in the database
	currentTime := time.Now()
	entity := MasterUrl{
		Key:         key,
		TargetUrl:   targetUrl,
		CreatedDate: &currentTime,
	}

	// Insert new record in the database
	error := repository.DB.Create(&entity).Error
	// Return error if any
	if error != nil {
		return nil, error
	}

	// Return the created record
	return &entity, nil
}
