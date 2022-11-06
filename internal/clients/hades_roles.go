package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
	"github.com/mondracode/ambrosia-atlas-api/internal/responses"
)

type HadesRoles struct {
	baseURL string
}

func NewHadesRoles(baseURL string) *HadesRoles {
	return &HadesRoles{
		baseURL: baseURL,
	}
}

func (hr *HadesRoles) GetUserRoles(userID string) (*responses.Roles, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/roles/user/%s", hr.baseURL, userID), nil)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

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

	if res.StatusCode == http.StatusNotFound {
		return nil, apperrors.NewNotFoundAppError("Roles in User", userID, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	var userRolesResponse responses.Roles

	err = UnmarshalBody(res.Body, &userRolesResponse)
	if err != nil {
		return nil, err
	}

	return &userRolesResponse, nil
}

func (hr *HadesRoles) GetUserScopes(userID string) (*responses.Scopes, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/scopes/user/%s", hr.baseURL, userID), nil)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

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

	if res.StatusCode == http.StatusNotFound {
		return nil, apperrors.NewNotFoundAppError("Scopes in User", userID, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	var userScopesResponse responses.Scopes

	err = UnmarshalBody(res.Body, &userScopesResponse)
	if err != nil {
		return nil, err
	}

	return &userScopesResponse, nil
}

func (hr *HadesRoles) GetRoleScopes(role string) (*[]string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/scopes/role/name/%s", hr.baseURL, role), nil)
	if err != nil {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

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

	if res.StatusCode == http.StatusNotFound {
		return nil, apperrors.NewNotFoundAppError("Scopes in Role", role, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, apperrors.NewUnexpectedAppError(err)
	}

	var roleScopesResponse responses.Scopes

	err = UnmarshalBody(res.Body, &roleScopesResponse)
	if err != nil {
		return nil, err
	}

	return roleScopesResponse.Scopes, nil
}
