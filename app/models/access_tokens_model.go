package models

import "time"

type OauthAcessTokens struct {
	UserId      string    `json:"user_id"`
	AccessToken string    `json:"access_token"`
	Expiry      time.Time `json:"expiry"`
	IsRevoke    bool      `json:"is_revoked"`
}
