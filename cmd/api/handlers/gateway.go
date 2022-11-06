package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/services"
)

type Gateway struct {
	allClients *clients.All
}

func NewGateway(allClients *clients.All) *Gateway {
	return &Gateway{allClients: allClients}
}

func (handler *Gateway) UseGateway(c *gin.Context) (interface{}, error) {
	graphQLQuery := &requests.GraphQLQuery{}

	err := c.ShouldBind(&graphQLQuery)
	if err != nil {
		return nil, apperrors.NewBadRequestAppError(err)
	}

	res, err := services.UseGateway(graphQLQuery, handler.allClients)
	if err != nil {
		return nil, err
	}

	return res, nil
}
