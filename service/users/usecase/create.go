package usecase

import (
	"sleekflow/helper"
	"sleekflow/models"
	"sleekflow/service/users/delivery/request"
	"sleekflow/utils/constants"
	"time"
)

func (u *Usecase) Create(request *request.UserCreateRequest) (*models.User, error) {

	pwd, _ := helper.HashPassword(request.Password)
	user, err := u.usersRepo.CreateUser(&models.User{
		Username:  request.Username,
		Password:  pwd,
		CreatedAt: time.Now().In(constants.TimeLocation),
	}, nil)

	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}
