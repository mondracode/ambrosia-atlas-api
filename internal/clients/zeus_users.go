package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

type ZeusUsers struct {
	baseURL string
}

func NewZeusUsers(baseURL string) *ZeusUsers {
	return &ZeusUsers{
		baseURL: baseURL,
	}
}

func (zu *ZeusUsers) Login(loginRequest *requests.Login) (*responses.ZeusLogin, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/users/login", zu.baseURL), requests.ToJSONBuffer(loginRequest))
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

	if res.StatusCode == http.StatusUnauthorized {
		return nil, apperrors.NewUnauthorizedAppError(loginRequest.Username, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	var loginResponse responses.ZeusLogin

	err = UnmarshalBody(res.Body, &loginResponse)
	if err != nil {
		return nil, err
	}

	return &loginResponse, err
}
