package users

import (
	"sleekflow/models"
	"sleekflow/service/users/delivery/request"
	"sleekflow/service/users/delivery/response"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	Login(request *request.LoginRequest, c *gin.Context) (*response.LoginResponse, error)
	Create(request *request.UserCreateRequest) (*models.User, error)
}
