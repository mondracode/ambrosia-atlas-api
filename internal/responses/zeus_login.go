package responses

type ZeusLogin struct {
	UserID      *string `json:"user_id"`
	Username    *string `json:"username"`
	DisplayName *string `json:"display_name"`
}
