package builder

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func InitClients() *clients.All {
	zeusUsersBaseURL := goDotEnvVariable("ZEUS_USERS_BASE_URL")
	hadesRolesBaseURL := goDotEnvVariable("HADES_ROLES_BASE_URL")
	cronosGatewayBaseURL := goDotEnvVariable("CRONOS_GATEWAY_BASE_URL")

	authJWTPassword := goDotEnvVariable("AUTH_JWT_PASSWORD")

	return &clients.All{
		ZeusUsers:     clients.NewZeusUsers(zeusUsersBaseURL),
		HadesRoles:    clients.NewHadesRoles(hadesRolesBaseURL),
		CronosGateway: clients.NewCronosGateway(cronosGatewayBaseURL),

		AuthClient: clients.NewAuthClient(authJWTPassword),
	}
}
