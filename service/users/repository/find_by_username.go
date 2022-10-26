package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) FindByUsername(username string) (*models.User, error) {
	model := models.User{}
	err := r.db.
		Where("username = ?", username).
		First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-user-bu-username:", err)
		return nil, errors.ErrUnprocessableEntity
	}

	return &model, nil
}
