package services

import (
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

func Login(loginRequest *requests.Login, allClients *clients.All) (*responses.CompleteLogin, error) {
	loginInfo, err := allClients.ZeusUsers.Login(loginRequest)
	if err != nil {
		return nil, err
	}

	roles, err := allClients.HadesRoles.GetUserRoles(*loginInfo.UserID)
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
