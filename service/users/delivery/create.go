package request

import (
	"log"
	"net/http"
	"sleekflow/service/users/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var request *request.UserCreateRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println("error-on-create-new-inventory: ", err)
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	userM, err := h.usersUsecase.Create(request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userM)
}
