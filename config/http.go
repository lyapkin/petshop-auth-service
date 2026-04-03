package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type HTTPServer struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (http *HTTPServer) Addr() string {
	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}

func loadHTTPConfig() (*HTTPServer, error) {
	host := os.Getenv("HTTP_HOST")
	if host == "" {
		return nil, errors.New("HTTP_HOST environment variable not set")
	}

	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("HTTP_PORT environment variable not integer: %w", err)
	}
	if port == 0 {
		return nil, errors.New("HTTP_PORT environment variable not set")
	}

	readTimeout, err := time.ParseDuration(os.Getenv("HTTP_READ_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("HTTP_WRITE_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	return &HTTPServer{
		Host:         host,
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}, nil
}
