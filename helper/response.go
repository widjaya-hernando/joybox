package helper

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BaseResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func NewResponse(status int, data interface{}, errors error) *BaseResponse {
	if errors != nil && !isValidationError(errors) {
		return &BaseResponse{
			Status: status,
			Data:   data,
			Errors: map[string]interface{}{"message": errors.Error()},
		}
	}

	return &BaseResponse{
		Status: status,
		Data:   data,
		Errors: errors,
	}
}

func isValidationError(err error) bool {
	_, ok := err.(validation.Errors)
	return ok
}
