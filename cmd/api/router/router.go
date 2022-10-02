package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/cmd/api/handler"
)

func SetupEndpoints() {
	router := gin.Default()

	router.GET("/ping", handler.Ping{}.Ping)

	err := router.Run()
	if err != nil {
		return
	}
}
