package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/requests"
)

type ZeusUsers struct {
	baseURL string
}

func NewZeusUsers(baseURL string) *ZeusUsers {
	return &ZeusUsers{
		baseURL: baseURL,
	}
}

func (zu *ZeusUsers) Login(loginRequest *requests.Login) (*string, error) {
	payload := strings.NewReader(
		fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"%s\"}",
			loginRequest.Username,
			loginRequest.Password,
		),
	)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/login", zu.baseURL), payload)
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

	var loginResponse LoginResponse

	err = UnmarshalBody(res.Body, &loginResponse)
	if err != nil {
		return nil, err
	}

	return loginResponse.UserID, err

}

type LoginResponse struct {
	UserID *string `json:"user_id"`
}
