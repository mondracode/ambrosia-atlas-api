package apperrors

import (
	"fmt"
	"net/http"
)

func NewUnauthorizedAppError(username string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, fmt.Sprintf("%s not authorized", username), err)
}
