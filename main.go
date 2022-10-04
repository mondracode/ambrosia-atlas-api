package main

import (
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/routers"
	"github.com/mondracode/ambrosia-atlas-api/internal/builder"
)

func main() {
	allClients := builder.InitClients()

	routers.SetupEndpoints(allClients)
}
