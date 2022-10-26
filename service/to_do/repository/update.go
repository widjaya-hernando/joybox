package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) Update(toDo *models.ToDo, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(toDo).Error
	if err != nil {
		log.Println("error-update-to-do:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
