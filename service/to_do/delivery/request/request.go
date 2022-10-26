package request

import (
	"sleekflow/lib/request_util"
	"time"
)

type ToDoCreateRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	Status      bool      `json:"status" binding:"required"`
}

type ToDoUpdateRequest struct {
	ID          uint64     `json:"id" binding:"required"`
	Name        *string    `json:"name" binding:"required"`
	Description *string    `json:"description" binding:"required"`
	DueDate     *time.Time `json:"due_date" binding:"required"`
	Status      *bool      `json:"status" binding:"required"`
}

type ToDoDeleteRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

func NewToDoPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	request_util.OverrideKey(conditions, "name", "to_do.name")
	request_util.OverrideKey(conditions, "description", "to_do.description")
	request_util.OverrideKey(conditions, "due_date", "to_do.due_date")
	request_util.OverrideKey(conditions, "status", "to_do.status")

	filterable := map[string]string{
		"to_do.sku":           request_util.StringType,
		"to_do.name":          request_util.StringType,
		"to_do.price":         request_util.DateType,
		"to_do.inventory_qty": request_util.BoolType,
	}
	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
