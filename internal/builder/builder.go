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
	zeusUsersBaseUrl := goDotEnvVariable("ZEUS_USERS_BASE_URL")

	return &clients.All{
		ZeusUsers: clients.NewZeusUsers(zeusUsersBaseUrl),
	}
}
