package usecase

import (
	"fmt"
	"sleekflow/helper"
	"sleekflow/service/users/delivery/request"
	"sleekflow/service/users/delivery/response"
	"sleekflow/utils/constants"
	"sleekflow/utils/errors"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) Login(request *request.LoginRequest, c *gin.Context) (*response.LoginResponse, error) {
	// validate if player exists
	user, err := u.usersRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.DoPasswordsMatch(user.Password, request.Password)
	if !isMatch {
		customError := errors.ErrUnauthorized
		customError.Message = "Invalid Username/Password"
		return nil, err
	}

	token, err := helper.GenerateToken(user, time.Now().In(constants.TimeLocation).Add(time.Hour*24), time.Now().In(constants.TimeLocation))
	if err != nil {
		customError := errors.ErrUnprocessableEntity
		customError.Message = fmt.Sprintf("Unable Generate Token, errors: %v", err)
		return nil, customError
	}

	return &response.LoginResponse{
		Token: token,
	}, nil
}
