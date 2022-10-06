package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/handlers"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/wrappers"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
)

func SetupEndpoints(allClients *clients.All) {
	router := gin.Default()

	errW := wrappers.NewErrorWrapper()
	authW := wrappers.NewAuthWrapper(allClients)

	ping := handlers.NewPing()
	router.GET("/ping", errW.Wrapper(ping.Ping))
	router.GET("/auth/ping", errW.Wrapper(authW.Wrapper(ping.Ping, "69_NICE")))

	session := handlers.NewSession(allClients)
	router.POST("/login", errW.Wrapper(session.Login))

	err := router.Run()
	if err != nil {
		return
	}

}
