package usecase

import (
	"log"
	"sleekflow/models"
	"sleekflow/service/to_do/delivery/request"
	"sleekflow/utils/errors"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) Create(request *request.ToDoCreateRequest, c *gin.Context) (*models.ToDo, error) {
	userID, valid := c.Get("user_id")
	if !valid {
		customError := errors.ErrFailedAuthentication
		customError.Message = "Unable to find User ID in token"
		return nil, customError
	}

	toDoM := &models.ToDo{
		UserID:      userID.(uint64),
		Name:        request.Name,
		Description: request.Description,
		DueDate:     request.DueDate,
		Status:      request.Status,
	}
	err := u.toDoRepo.Create(toDoM, nil)
	if err != nil {
		log.Println("error-on-create-new-inventory: ", err)
		return nil, err
	}

	return toDoM, nil
}
