package clients

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jtblin/go-ldap-client"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/models"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

type AuthClient struct {
	*ldap.LDAPClient
	jwtPassword []byte
}

func NewLDAPClient(authLDAPConfig map[string]string) *ldap.LDAPClient {
	port, _ := strconv.Atoi(authLDAPConfig["port"])
	useSSL, _ := strconv.ParseBool(authLDAPConfig["useSSL"])
	skipTLS, _ := strconv.ParseBool(authLDAPConfig["skipTLS"])

	return &ldap.LDAPClient{
		Base:         authLDAPConfig["base"],
		Host:         authLDAPConfig["host"],
		Port:         port,
		UseSSL:       useSSL,
		SkipTLS:      skipTLS,
		BindDN:       authLDAPConfig["bindDN"],
		BindPassword: authLDAPConfig["bindPassword"],
		UserFilter:   authLDAPConfig["userFilter"],
		GroupFilter:  authLDAPConfig["groupFilter"],
		Attributes:   []string{"givenName", "sn", "mail", "uid"},
		ServerName:   authLDAPConfig["serverName"],
	}
}

func NewAuthClient(authLDAPconfig map[string]string, jwtPassword string) *AuthClient {
	return &AuthClient{
		LDAPClient:  NewLDAPClient(authLDAPconfig),
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

func (a *AuthClient) ValidateJWT(bearerJWT string) error {
	token, err := jwt.ParseWithClaims(bearerJWT, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.jwtPassword, nil
	})

	if err != nil {
		return apperrors.NewBadRequestAppError(err)
	}

	claims, ok := token.Claims.(*models.JWTClaims)

	if !(ok && token.Valid) {
		return apperrors.NewUnauthorizedAppError(*claims.ZeusLogin.Username, err)
	}

	return nil
}
