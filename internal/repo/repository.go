package repo

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewReposytory(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
