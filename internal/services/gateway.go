package services

import (
	"github.com/mondracode/ambrosia-atlas-api/internal/clients"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

func UseGateway(graphQLQuery *requests.GraphQLQuery, allClients *clients.All) (*responses.Gateway, error) {
	res, err := allClients.CronosGateway.UseGateway(graphQLQuery)
	if err != nil {
		return nil, err
	}

	return res, nil
}
