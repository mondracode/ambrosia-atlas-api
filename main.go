package main

import (
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/router"
	"github.com/mondracode/ambrosia-atlas-api/internal/builder"
)

func main() {
	allClients := builder.InitClients()

	router.SetupEndpoints(allClients)
}
