package jwttoken

import (
	"crypto/rsa"
	"time"

	"github.com/lyapkin/shop/auth/config"
)

type service struct {
	accessSecret *rsa.PrivateKey
	accessTTL    time.Duration
	refreshTTL   time.Duration
}

func New(cfg *config.JWTToken) *service {
	return &service{
		accessSecret: cfg.AccessSecret,
		accessTTL:    cfg.AccessTTL,
		refreshTTL:   cfg.RefreshTTL,
	}
}
