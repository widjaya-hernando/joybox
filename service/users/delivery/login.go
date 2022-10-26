package request

import (
	"net/http"
	"sleekflow/helper"
	"sleekflow/service/users/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := h.usersUsecase.Login(&request, c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}
