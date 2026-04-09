package config

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	AccessSecret *rsa.PrivateKey
	AccessPublic *rsa.PublicKey
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
}

func loadJWTTokenConfig() (*JWTToken, error) {
	jwtAccessPrivatePath := os.ExpandEnv(os.Getenv("JWT_ACCESS_SECRET_PATH"))
	if jwtAccessPrivatePath == "" {
		return nil, errors.New("no JWT_ACCESS_SECRET_PATH environment variable")
	}

	jwtAccessPublicPath := os.ExpandEnv(os.Getenv("JWT_ACCESS_PUBLIC_PATH"))
	if jwtAccessPublicPath == "" {
		return nil, errors.New("no JWT_ACCESS_PUBLIC_PATH environment variable")
	}

	jwtAccessTTL, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	if err != nil {
		return nil, err
	}

	jwtRefreshTTL, err := time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))
	if err != nil {
		return nil, err
	}

	jwtAccessPrivate, err := loadRSAPrivateKey(jwtAccessPrivatePath)
	if err != nil {
		return nil, err
	}

	jwtAccessPublic, err := loadRSAPublicKey(jwtAccessPublicPath)
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessSecret: jwtAccessPrivate,
		AccessPublic: jwtAccessPublic,
		AccessTTL:    jwtAccessTTL,
		RefreshTTL:   jwtRefreshTTL,
	}, nil
}

func loadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
}

func loadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyBytes)
}
