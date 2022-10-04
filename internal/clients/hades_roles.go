package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
)

type HadesRoles struct {
	baseURL string
}

func NewHadesRoles(baseURL string) *HadesRoles {
	return &HadesRoles{
		baseURL: baseURL,
	}
}

func (hr *HadesRoles) GetAllUserScopes(userID string) (*[]string, error) {
	scopes, err := hr.GetUserScopes(userID)
	if err != nil {
		scopes = &[]string{}
	}

	roles, err := hr.GetUserRoles(userID)
	if err != nil {
		roles = &[]string{}
	}

	for _, role := range *roles {
		roleScopes, _ := hr.GetRoleScopes(role)
		*scopes = append(*scopes, *roleScopes...)
	}

	if len(*scopes) == 0 {
		return nil, apperrors.NewNotFoundAppError("Scopes in User", userID, nil)
	}

	return scopes, nil
}

func (hr *HadesRoles) GetUserRoles(userID string) (*[]string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles/user/%s", hr.baseURL, userID), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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

	var userRolesResponse RolesResponse

	err = UnmarshalBody(res.Body, &userRolesResponse)
	if err != nil {
		return nil, err
	}

	return userRolesResponse.Roles, nil
}

func (hr *HadesRoles) GetUserScopes(userID string) (*[]string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/scopes/user/%s", hr.baseURL, userID), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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

	var userScopesResponse ScopesResponse

	err = UnmarshalBody(res.Body, &userScopesResponse)
	if err != nil {
		return nil, err
	}

	return userScopesResponse.Scopes, nil
}

func (hr *HadesRoles) GetRoleScopes(role string) (*[]string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/scopes/role/name/%s", hr.baseURL, role), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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

	var roleScopesResponse ScopesResponse

	err = UnmarshalBody(res.Body, &roleScopesResponse)
	if err != nil {
		return nil, err
	}

	return roleScopesResponse.Scopes, nil
}

type RolesResponse struct {
	Roles *[]string `json:"roles"`
}

type ScopesResponse struct {
	Scopes *[]string `json:"scopes"`
}
