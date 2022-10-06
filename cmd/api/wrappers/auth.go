package wrappers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
)

type AuthWrapper struct {
	AllClients *clients.All
}

func NewAuthWrapper(allClients *clients.All) *AuthWrapper {
	return &AuthWrapper{AllClients: allClients}
}

func (a AuthWrapper) Wrapper(f func(c *gin.Context) (interface{}, error), scopeRequired string) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return nil, apperrors.NewBadRequestAppError(fmt.Errorf("invalid authorization header"))
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

		err := a.AllClients.AuthClient.ValidateJWT(bearerToken, scopeRequired)
		if err != nil {
			return nil, err
		}

		return f(c)
	}
}
