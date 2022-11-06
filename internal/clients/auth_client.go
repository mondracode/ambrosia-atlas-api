package clients

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/models"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
	"golang.org/x/exp/slices"
)

type AuthClient struct {
	jwtPassword []byte
}

func NewAuthClient(jwtPassword string) *AuthClient {
	return &AuthClient{
		jwtPassword: []byte(jwtPassword),
	}
}

func (a *AuthClient) GenerateJWT(loginInfo responses.ZeusLogin, roles, scopes []string) (*string, error) {

	claims := models.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		ZeusLogin: loginInfo,
		Roles:     roles,
		Scopes:    scopes,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(a.jwtPassword)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	return &tokenStr, nil
}

func (a *AuthClient) ValidateJWT(bearerJWT, scopeRequired string) error {
	token, err := jwt.ParseWithClaims(bearerJWT, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.jwtPassword, nil
	})

	if err != nil {
		return apperrors.NewBadRequestAppError(err)
	}

	claims, ok := token.Claims.(*models.JWTClaims)

	hasScopeRequired := slices.Contains(claims.Scopes, strings.ToUpper(scopeRequired))

	if !(ok && token.Valid && hasScopeRequired) {
		return apperrors.NewUnauthorizedAppError(*claims.ZeusLogin.Username, err)
	}

	return nil
}
