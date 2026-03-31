package domain

import "time"

type Token struct {
	AccessToken
	RefreshToken
}

type RefreshToken struct {
	Token     string    `json:"refreshToken"`
	ExpiresAt time.Time `json:"-"`
}

type AccessToken struct {
	Token string `json:"accessToken"`
}
