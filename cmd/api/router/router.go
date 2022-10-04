package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/handler"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
)

func SetupEndpoints(AllClients *clients.All) {
	router := gin.Default()

	router.GET("/ping", handler.Ping{}.Ping)

	session := handler.NewSession(AllClients.ZeusUsers)
	router.POST("/login", session.Login)

	err := router.Run()
	if err != nil {
		return
	}
}
