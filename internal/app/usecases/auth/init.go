package auth

import (
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type usecase struct {
	log          *slog.Logger
	userRepo     interfaces.UserRepo
	roleRepo     interfaces.RoleRepo
	password     interfaces.PasswordHasher
	tokenService interfaces.TokenService
	tokenRepo    interfaces.TokenRepo
}

func New(
	log *slog.Logger,
	userRepo interfaces.UserRepo,
	roleRepo interfaces.RoleRepo,
	password interfaces.PasswordHasher,
	tokenService interfaces.TokenService,
	tokenRepo interfaces.TokenRepo,
) *usecase {
	return &usecase{
		log:          log,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		password:     password,
		tokenService: tokenService,
		tokenRepo:    tokenRepo,
	}
}
