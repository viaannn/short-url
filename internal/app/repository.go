package app

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
