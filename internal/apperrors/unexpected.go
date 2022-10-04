package apperrors

import (
	"net/http"
)

func NewUnexpectedAppError(err error) *AppError {
	return &AppError{
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        "Unknown error",
		WrappedError:   err,
	}
}
