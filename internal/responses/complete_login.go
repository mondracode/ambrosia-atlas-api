package responses

type CompleteLogin struct {
	*ZeusLogin
	*Roles
	*Scopes
	AuthToken *string `json:"auth_token"`
}
