package usecase

import (
	"sleekflow/models"
	"sleekflow/service/to_do/delivery/request"
	"sleekflow/utils/errors"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) Delete(request *request.ToDoDeleteRequest, c *gin.Context) error {
	var toDoM *models.ToDo
	var err error

	userID, valid := c.Get("user_id")
	if !valid {
		customError := errors.ErrFailedAuthentication
		customError.Message = "Unable to find User ID in token"
		return customError
	}

	tx := u.transactionManager.NewTransaction()
	tx.Begin()
	{
		toDoM, err = u.toDoRepo.FindByID(request.ID)
		if err != nil {
			tx.Rollback()
			return err
		}

		if userID.(uint64) != toDoM.UserID {
			customError := errors.ErrUnauthorized
			customError.Message = "You are unauthorized to access this to-do"
			return customError
		}

		err = u.toDoRepo.Delete(toDoM, tx)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}
