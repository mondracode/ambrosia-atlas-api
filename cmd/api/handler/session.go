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
	zeusUsersClient *clients.ZeusUsers
}

func NewSession(zeusUsersClient *clients.ZeusUsers) *Session {
	return &Session{zeusUsersClient: zeusUsersClient}
}

func (handler *Session) Login(ctx *gin.Context) {
	login := &requests.Login{}

	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	err = services.Login(login, handler.zeusUsersClient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusOK, login)
	return
}
