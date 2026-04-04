package domain

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	AccessToken
	RefreshToken
}

type RefreshToken struct {
	Token     string    `json:"refreshToken"`
	ExpiresAt time.Time `json:"-"`
	UserID    uuid.UUID `json:"-"`
}

type AccessToken struct {
	Token string `json:"accessToken"`
}
