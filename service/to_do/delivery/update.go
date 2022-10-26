package request

import (
	"net/http"

	"sleekflow/service/to_do/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	var request *request.ToDoUpdateRequest

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	toDoM, err := h.toDoUsecase.Update(request, c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, toDoM)
}
