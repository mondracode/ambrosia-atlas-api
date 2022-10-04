package clients

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
)

type AuthClient struct {
	jwtPassword []byte
}

func NewAuthClient(jwtPassword string) *AuthClient {
	return &AuthClient{
		jwtPassword: []byte(jwtPassword),
	}
}

func (a *AuthClient) GenerateJWT(userID, username string, scopes []string) (string, error) {

	payload := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"scopes":   scopes,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(a.jwtPassword)
	if err != nil {
		return tokenStr, apperrors.NewUnexpectedAppError(err)
	}

	return tokenStr, nil
}
