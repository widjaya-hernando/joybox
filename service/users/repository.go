package users

import (
	"sleekflow/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *models.User, tx *gorm.DB) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}
