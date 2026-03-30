package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DBDriver string

const (
	POSTGRES DBDriver = "postgres"
)

type DB struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
	Driver DBDriver
}

func (db *DB) URL() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		db.Driver,
		db.User,
		db.Pass,
		db.Host,
		db.Port,
		db.Name,
	)
}

func parseDBDriver(s string) (DBDriver, error) {
	switch strings.ToLower(s) {
	case "postgres":
		return POSTGRES, nil
	case "":
		return "", errors.New("DB_DRIVER environment variable not set")
	default:
		return "", errors.New("unkown db driver")
	}
}

func loadDBConfig() (*DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("DB_HOST environment variable not set")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("DB_PORT environment variable not integer: %w", err)
	}
	if port == 0 {
		return nil, errors.New("DB_PORT environment variable not set")
	}

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		return nil, errors.New("DB_USERNAME environment variable not set")
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return nil, errors.New("DB_PASSWORD environment variable not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, errors.New("DB_NAME environment variable not set")
	}

	driver, err := parseDBDriver(os.Getenv("DB_DRIVER"))
	if err != nil {
		return nil, err
	}

	return &DB{
		Host:   host,
		Port:   port,
		User:   username,
		Pass:   pass,
		Name:   dbName,
		Driver: DBDriver(driver),
	}, nil
}
