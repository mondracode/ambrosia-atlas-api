package services

import (
	"fmt"

	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
)

func Login(loginRequest *requests.Login, zeusUsersClient *clients.ZeusUsers) error {
	userID, err := zeusUsersClient.Login(loginRequest)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("aaaaa userID: %v", *userID))
	return nil
}
