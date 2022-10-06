package wrappers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
)

type ErrorWrapper struct{}

func NewErrorWrapper() *ErrorWrapper {
	return &ErrorWrapper{}
}

func (e ErrorWrapper) Wrapper(f func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := f(c)
		if err != nil {
			var appError *apperrors.AppError
			switch {
			case errors.As(err, &appError):
				c.JSON(appError.HTTPStatusCode, gin.H{
					"message": appError.Error(),
					"error":   appError.ErrorType,
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}
