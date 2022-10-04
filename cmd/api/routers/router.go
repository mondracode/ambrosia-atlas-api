package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/handlers"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/wrappers"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
)

func SetupEndpoints(allClients *clients.All) {
	router := gin.Default()

	router.GET("/ping", handlers.Ping{}.Ping)

	session := handlers.NewSession(allClients)
	router.POST("/login", wrappers.ErrorWrapper(session.Login))

	err := router.Run()
	if err != nil {
		return
	}
}
