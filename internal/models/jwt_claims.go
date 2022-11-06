package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

type JWTClaims struct {
	responses.ZeusLogin
	jwt.RegisteredClaims
	Roles  []string `json:"roles"`
	Scopes []string `json:"scopes"`
}
