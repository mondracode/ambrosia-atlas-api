package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/services"
)

type Session struct {
	allClients *clients.All
}

func NewSession(allClients *clients.All) *Session {
	return &Session{allClients: allClients}
}

func (handler *Session) Login(ctx *gin.Context) {
	login := &requests.Login{}

	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	jwt, err := services.Login(login, handler.allClients)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusOK, jwt)
	return
}
