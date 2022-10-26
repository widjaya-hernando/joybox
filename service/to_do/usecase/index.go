package usecase

import (
	"fmt"

	"sleekflow/models"
	"sleekflow/utils/errors"

	"sleekflow/lib/request_util"
	"sleekflow/lib/response_util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u *Usecase) Index(paginationConfig request_util.PaginationConfig, c *gin.Context) ([]models.ToDo, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	userID, valid := c.Get("user_id")
	if !valid {
		customError := errors.ErrFailedAuthentication
		customError.Message = "Unable to find User ID in token"
		return nil, meta, customError
	}

	paginationConfig.AddScope(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID.(uint64))
	})

	if search, ok := c.Request.URL.Query()["search"]; ok {
		paginationConfig.AddScope(func(db *gorm.DB) *gorm.DB {
			return db.Where("AND sku like ? OR name like ? ", fmt.Sprint("%", search[0], "%"), fmt.Sprint("%", search[0], "%"))
		})
	}

	inventories, err := u.toDoRepo.FindAll(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	total, err := u.toDoRepo.Count(paginationConfig)
	if err != nil {
		return nil, meta, err
	}
	meta.Total = total

	return inventories, meta, nil
}
