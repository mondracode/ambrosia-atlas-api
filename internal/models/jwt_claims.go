package models

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Scopes   []string `json:"scopes"`
}
