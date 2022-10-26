package repository

import (
	"log"

	"sleekflow/lib/request_util"
	"sleekflow/models"
	"sleekflow/utils/errors"
)

func (r *Repository) Count(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.ToDo{}).
		Scopes(config.Scopes()...).
		Count(&count).Error

	if err != nil {
		log.Println("error-count-to-do:", err)
		return 0, errors.ErrUnprocessableEntity
	}
	return count, nil
}
