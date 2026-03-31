package jwttoken

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/config"
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

func New(cfg *config.JWTToken) *service {
	return &service{
		accessSecret:  cfg.AccessSecret,
		refreshSecret: cfg.RefreshSecret,
		accessTTL:     cfg.AccessTTL,
		refreshTTL:    cfg.RefreshTTL,
	}
}
