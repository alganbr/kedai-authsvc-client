package models

type AccessToken struct {
	Id      string `json:"id"`
	UserId  int64  `json:"user_id"`
	Expires int64  `json:"expires"`
}

type AccessTokenRq struct {
	GrantType string `json:"grant_type"`

	// Grant type password
	Email    string `json:"email"`
	Password string `json:"password"`

	// Grant type oauth
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
