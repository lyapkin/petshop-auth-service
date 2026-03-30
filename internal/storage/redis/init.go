package redis

import (
	"context"

	"github.com/lyapkin/shop/auth/config"
	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, cfg config.Redis) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Username: cfg.Host,
		Password: cfg.Pass,
	})

	err := db.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return db, nil
}
