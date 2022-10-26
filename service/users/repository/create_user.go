package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) CreateUser(user *models.User, tx *gorm.DB) (*models.User, error) {
	var db = r.db

	if tx != nil {
		db = tx
	}
	err := db.Create(user).Error
	if err != nil {
		log.Println("error-insert-user:", err)
		return nil, errors.ErrUnprocessableEntity
	}
	return user, nil
}
