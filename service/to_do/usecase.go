package to_do

import (
	"sleekflow/lib/request_util"
	"sleekflow/lib/response_util"
	"sleekflow/models"
	"sleekflow/service/to_do/delivery/request"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	Index(paginationConfig request_util.PaginationConfig, c *gin.Context) ([]models.ToDo, response_util.PaginationMeta, error)
	Show(c *gin.Context) (*models.ToDo, error)
	Create(request *request.ToDoCreateRequest, c *gin.Context) (*models.ToDo, error)
	Update(request *request.ToDoUpdateRequest, c *gin.Context) (*models.ToDo, error)
	Delete(request *request.ToDoDeleteRequest, c *gin.Context) error
}
