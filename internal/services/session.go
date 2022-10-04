package services

import (
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
)

func Login(loginRequest *requests.Login, allClients *clients.All) (*string, error) {
	userID, err := allClients.ZeusUsers.Login(loginRequest)
	if err != nil {
		return nil, err
	}

	scopes, err := allClients.HadesRoles.GetUserScopes(*userID)
	if err != nil {
		scopes = &[]string{}
	}

	jwt, err := allClients.AuthClient.GenerateJWT(*userID, loginRequest.Username, *scopes)

	return &jwt, err
}
