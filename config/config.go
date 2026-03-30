package config

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type config struct {
	Env Env
	JWTToken
	DB
	Redis
}

type JWTToken struct {
	AccessSecret  *rsa.PrivateKey
	AccessPublic  *rsa.PublicKey
	RefreshSecret []byte
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

func MustLoad() *config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	env := ParseEnv(os.Getenv("ENV"))

	jwtToken, err := loadJWTTokenConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := loadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	redis, err := loadRedisConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &config{
		Env:      env,
		JWTToken: *jwtToken,
		DB:       *db,
		Redis:    *redis,
	}
}

func loadJWTTokenConfig() (*JWTToken, error) {
	jwtAccessPrivatePath := os.Getenv("JWT_ACCESS_SECRET_PATH")
	if jwtAccessPrivatePath == "" {
		return nil, errors.New("no JWT_ACCESS_SECRET_PATH environment variable")
	}

	jwtAccessPublicPath := os.Getenv("JWT_ACCESS_PUBLIC_PATH")
	if jwtAccessPublicPath == "" {
		return nil, errors.New("no JWT_ACCESS_PUBLIC_PATH environment variable")
	}

	jwtRefreshPath := os.Getenv("JWT_REFRESH_SECRET_PATH")
	if jwtRefreshPath == "" {
		return nil, errors.New("no JWT_REFRESH_SECRET_PATH environment variable")
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

	jwtRefresh, err := loadStringKey(jwtRefreshPath)
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessSecret:  jwtAccessPrivate,
		AccessPublic:  jwtAccessPublic,
		RefreshSecret: jwtRefresh,
		AccessTTL:     jwtAccessTTL,
		RefreshTTL:    jwtRefreshTTL,
	}, nil
}

func loadStringKey(path string) ([]byte, error) {
	key, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(string(key))), nil
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
