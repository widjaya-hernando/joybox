package request

import (
	"log"
	"net/http"
	"sleekflow/service/to_do/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	var request *request.ToDoDeleteRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println("error-on-create-new-inventory: ", err)
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	err := h.toDoUsecase.Delete(request, c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
