package request

import (
	"log"
	"net/http"
	"sleekflow/service/to_do/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var request *request.ToDoCreateRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println("error-on-create-new-inventory: ", err)
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	inventoryM, err := h.toDoUsecase.Create(request, c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, inventoryM)
}
