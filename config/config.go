package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	Env              Env
	ShutdownTimeout  time.Duration
	InMemoryCacheTTL time.Duration
	JWTToken
	DB
	Redis
	HTTPServer
}

func MustLoad() *config {
	if err := godotenv.Load(); err != nil {
		log.Printf("%v using system environment variables", err)
	}

	env := ParseEnv(os.Getenv("ENV"))

	shutdownTimeout, err := time.ParseDuration(os.Getenv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		log.Fatalf("Can not parse SHUTDOWN_TIMEOUT: %v", err)
	}

	inMemoryCacheTTL, err := time.ParseDuration(os.Getenv("IN_MEMROY_CACHE_TTL"))
	if err != nil {
		log.Fatalf("Can not parse IN_MEMROY_CACHE_TTL: %v", err)
	}

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

	http, err := loadHTTPConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &config{
		Env:              env,
		ShutdownTimeout:  shutdownTimeout,
		InMemoryCacheTTL: inMemoryCacheTTL,
		JWTToken:         *jwtToken,
		DB:               *db,
		Redis:            *redis,
		HTTPServer:       *http,
	}
}
