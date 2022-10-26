package request

import (
	"net/http"

	"sleekflow/lib/response_util"
	"sleekflow/service/to_do/delivery/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	inventories, inventoryPagination, err := h.toDoUsecase.Index(request.NewToDoPaginationConfig(c.Request.URL.Query()), c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Data: inventories,
		Meta: inventoryPagination,
	})
}
