package apperrors

import (
	"fmt"
	"net/http"
)

func NewUnauthorizedAppError(username string, err error) *AppError {
	return &AppError{
		HTTPStatusCode: http.StatusUnauthorized,
		Message:        fmt.Sprintf("%s not authorized", username),
		WrappedError:   err,
	}
}
