package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Redis struct {
	Host string
	Port int
	User string
	Pass string
}

func (db *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", db.Host, db.Port)
}

func loadRedisConfig() (*Redis, error) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return nil, errors.New("REDIS_HOST environment variable not set")
	}

	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, fmt.Errorf("REDIS_PORT environment variable not integer: %w", err)
	}
	if port == 0 {
		return nil, errors.New("REDIS_PORT environment variable not set")
	}

	username := os.Getenv("REDIS_USERNAME")
	if username == "" {
		return nil, errors.New("REDIS_USERNAME environment variable not set")
	}

	pass := os.Getenv("REDIS_PASSWORD")
	if pass == "" {
		return nil, errors.New("REDIS_PASSWORD environment variable not set")
	}

	return &Redis{
		Host: host,
		Port: port,
		User: username,
		Pass: pass,
	}, nil
}
