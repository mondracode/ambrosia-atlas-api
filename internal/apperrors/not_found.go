package apperrors

import (
	"fmt"
	"net/http"
)

func NewNotFoundAppError(entityName string, entityID string, err error) *AppError {
	return &AppError{
		HTTPStatusCode: http.StatusNotFound,
		Message:        fmt.Sprintf("%s %s Not found", entityName, entityID),
		WrappedError:   err,
	}
}
