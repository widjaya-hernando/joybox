package repository

import (
	"log"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"sleekflow/lib/request_util"
)

func (r *Repository) FindAll(config request_util.PaginationConfig) ([]models.ToDo, error) {
	var results []models.ToDo
	err := r.db.
		Scopes(config.Scopes()...).
		Find(&results).Debug().Error
	if err != nil {
		log.Println("error-find-all-to-do:", err)
		return nil, errors.ErrUnprocessableEntity
	}
	return results, nil
}
