package clients

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

type CronosGateway struct {
	baseURL string
}

func NewCronosGateway(baseURL string) *CronosGateway {
	return &CronosGateway{
		baseURL: baseURL,
	}
}

func (cg *CronosGateway) UseGateway(graphQLQuery *requests.GraphQLQuery) (*responses.Gateway, error) {
	req, err := http.NewRequest(http.MethodPost, cg.baseURL+"/graphql", requests.ToJSONBuffer(graphQLQuery))
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error in body client")
		}
	}(res.Body)

	var gatewayResponse responses.Gateway

	err = UnmarshalBody(res.Body, &gatewayResponse)
	if err != nil {
		return nil, err
	}

	if gatewayResponse.Errors != nil {
		return nil, apperrors.NewUnexpectedAppError(errors.New(gatewayResponse.ErrorsToString()))
	}

	return &gatewayResponse, nil
}
