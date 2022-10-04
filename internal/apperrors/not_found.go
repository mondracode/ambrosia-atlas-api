package apperrors

import (
	"fmt"
	"net/http"
)

func NewNotFoundAppError(entityName string, entityID string, err error) *AppError {
	return NewAppError(http.StatusNotFound, fmt.Sprintf("%s %s Not found", entityName, entityID), err)
}
