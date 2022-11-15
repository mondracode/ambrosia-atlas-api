package services

import (
	"strings"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

func Login(loginRequest *requests.Login, allClients *clients.All) (*responses.CompleteLogin, error) {
	// Login with LDAP

	ok, user, err := allClients.AuthClient.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}
	if !ok {
		return nil, apperrors.NewUnauthorizedAppError(loginRequest.Username, err)
	}

	// Login with Zeus and Hades

	loginInfo, err := allClients.ZeusUsers.Login(loginRequest)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(user["uid"], *loginInfo.Username) {
		return nil, apperrors.NewUnauthorizedAppError(loginRequest.Username, err)
	}

	roles, err := allClients.HadesRoles.GetUserRoles(*loginInfo.UserID)
	if err != nil {
		roles = &responses.Roles{}
	}
	scopes, err := allClients.HadesRoles.GetUserScopes(*loginInfo.UserID)
	if err != nil {
		scopes = &responses.Scopes{}
	}

	jwt, err := allClients.AuthClient.GenerateJWT(*loginInfo, *roles.Roles, *scopes.Scopes)
	if err != nil {
		return nil, err
	}

	completeLogin := responses.CompleteLogin{
		ZeusLogin: loginInfo,
		Roles:     roles,
		Scopes:    scopes,
		AuthToken: jwt,
	}

	return &completeLogin, nil
}
