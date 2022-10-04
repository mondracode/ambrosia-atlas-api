package apperrors

import (
	"net/http"
)

func NewUnexpectedAppError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "Unknown error", err)
}
