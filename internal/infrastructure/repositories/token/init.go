package token

import "github.com/redis/go-redis/v9"

var refreshTokenKeyPrefix = "refreshToken:"

type tokenRepo struct {
	db *redis.Client
}

func New(db *redis.Client) *tokenRepo {
	return &tokenRepo{
		db: db,
	}
}
