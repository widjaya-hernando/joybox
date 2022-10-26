package request

import (
	"sleekflow/app/api/middleware"
	"sleekflow/service/users"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usersUsecase users.Usecase
}

func New(usersUC users.Usecase) *Handler {
	return &Handler{
		usersUsecase: usersUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	usersRoute := r.Group("/users")
	{
		usersRoute.POST("/login", h.Login)
		usersRoute.POST("/create", h.Create)
	}
}
