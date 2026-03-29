package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type refreshTokenClaims struct {
	TokenID uuid.UUID `json:"jti"`
	jwt.RegisteredClaims
}

type service struct {
	accessSecret  *rsa.PrivateKey
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func New() *service {
	return &service{}
}
