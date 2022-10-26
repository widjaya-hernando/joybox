package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Show(c *gin.Context) {
	inventoryM, err := h.toDoUsecase.Show(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, inventoryM)
}
