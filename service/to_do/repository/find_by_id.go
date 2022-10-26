package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) FindByID(ID uint64) (*models.ToDo, error) {
	model := models.ToDo{}
	err := r.db.
		Where("id = ?", ID).
		First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-to-do-by-id:", err)
		return nil, errors.ErrUnprocessableEntity
	}

	return &model, nil
}
