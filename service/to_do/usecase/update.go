package usecase

import (
	"log"
	"sleekflow/models"
	"sleekflow/service/to_do/delivery/request"
	"sleekflow/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func (u *Usecase) Update(request *request.ToDoUpdateRequest, c *gin.Context) (*models.ToDo, error) {
	userID, valid := c.Get("user_id")
	if !valid {
		customError := errors.ErrFailedAuthentication
		customError.Message = "Unable to find User ID in token"
		return nil, customError
	}

	toDo, err := u.toDoRepo.FindByID(request.ID)
	if err != nil {
		return nil, err
	}

	if userID.(uint64) != toDo.UserID {
		customError := errors.ErrUnauthorized
		customError.Message = "You are unauthorized to access this to-do"
		return nil, customError
	}

	err = copier.Copy(toDo, &request)
	if err != nil {
		log.Println("error-for-copy-request-to-inventory")
		return nil, err
	}

	err = u.toDoRepo.Update(toDo, nil)
	if err != nil {
		return nil, err
	}

	return toDo, err
}
