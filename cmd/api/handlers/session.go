package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
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

func (handler *Session) Login(ctx *gin.Context) (interface{}, error) {
	login := &requests.Login{}

	err := ctx.ShouldBind(&login)
	if err != nil {
		return nil, apperrors.NewBadRequestAppError(err)
	}

	jwt, err := services.Login(login, handler.allClients)
	if err != nil {
		return nil, err
	}

	return jwt, err
}
