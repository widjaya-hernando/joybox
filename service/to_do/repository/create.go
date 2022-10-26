package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) Create(toDo *models.ToDo, tx *gorm.DB) error {
	var db = r.db

	if tx != nil {
		db = tx
	}
	err := db.Debug().Create(toDo).Error
	if err != nil {
		log.Println("error-insert-to-do:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
