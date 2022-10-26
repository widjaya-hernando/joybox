package repository

import (
	"sleekflow/service/to_do"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) to_do.Repository {
	return &Repository{
		db: db,
	}
}
