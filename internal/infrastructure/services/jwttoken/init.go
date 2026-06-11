package jwttoken

import (
	"crypto/rsa"
	"time"

	"github.com/lyapkin/shop/auth/config"
)

type service struct {
	accessSecret *rsa.PrivateKey
	accessPublic *rsa.PublicKey
	accessTTL    time.Duration
	refreshTTL   time.Duration
}

func New(cfg *config.JWTToken) *service {
	return &service{
		accessSecret: cfg.AccessSecret,
		accessPublic: cfg.AccessPublic,
		accessTTL:    cfg.AccessTTL,
		refreshTTL:   cfg.RefreshTTL,
	}
}
