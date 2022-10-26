package repository

import (
	"sleekflow/models"
	"sleekflow/service/users"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Model struct {
	productM models.User
}

func New(
	db *gorm.DB,
) users.Repository {
	return &Repository{
		db: db,
	}
}
