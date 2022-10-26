package to_do

import (
	"sleekflow/lib/request_util"
	"sleekflow/models"

	"gorm.io/gorm"
)

type Repository interface {
	Delete(inventory *models.ToDo, tx *gorm.DB) error
	FindByID(ID uint64) (*models.ToDo, error)
	FindAll(config request_util.PaginationConfig) ([]models.ToDo, error)
	Create(inventory *models.ToDo, tx *gorm.DB) error
	Update(inventory *models.ToDo, tx *gorm.DB) error
	Count(config request_util.PaginationConfig) (int64, error)
}
