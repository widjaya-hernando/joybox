package usecase

import (
	"sleekflow/models"
	"sleekflow/utils/errors"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) Show(c *gin.Context) (*models.ToDo, error) {
	userID, valid := c.Get("user_id")
	if !valid {
		customError := errors.ErrFailedAuthentication
		customError.Message = "Unable to find User ID in token"
		return nil, customError
	}

	ID, valid := c.Get("id")
	if !valid {
		customError := errors.ErrBadRequest
		customError.Message = "id is required"
	}

	toDo, err := u.toDoRepo.FindByID(ID.(uint64))
	if err != nil {
		return nil, err
	}

	if userID.(uint64) != toDo.UserID {
		customError := errors.ErrUnauthorized
		customError.Message = "You are unauthorized to access this to-do"
		return nil, customError
	}

	return toDo, nil
}
