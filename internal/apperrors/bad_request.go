package apperrors

import (
	"net/http"
)

func NewBadRequestAppError(err error) *AppError {
	return NewAppError(http.StatusBadRequest, "Bad request", err)
}
