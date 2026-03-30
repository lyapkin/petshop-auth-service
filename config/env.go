package config

import "strings"

type Env string

const (
	EnvDevelopment Env = "DEVELOPMENT"
	EnvTest        Env = "TEST"
	EnvProduction  Env = "PRODUCTION"
)

func ParseEnv(s string) Env {
	switch strings.ToLower(s) {
	case "production", "prod":
		return EnvProduction
	case "test":
		return EnvTest
	default:
		return EnvDevelopment
	}
}
