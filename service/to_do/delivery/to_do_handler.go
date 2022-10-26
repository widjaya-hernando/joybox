package request

import (
	"sleekflow/app/api/middleware"
	"sleekflow/service/to_do"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	toDoUsecase to_do.Usecase
}

func New(toDoUC to_do.Usecase) *Handler {
	return &Handler{
		toDoUsecase: toDoUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	toDoRoute := r.Group("/to-do", m.Authorization)
	{
		toDoRoute.GET("/list", h.Index)
		toDoRoute.GET("/show", h.Show)
		toDoRoute.POST("/create", h.Create)
		toDoRoute.DELETE("/delete", h.Delete)
		toDoRoute.PUT("/update", h.Update)
	}
}
