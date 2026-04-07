package auth

import (
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type Usecase struct {
	log          *slog.Logger
	accountRepo  interfaces.AccountRepo
	roleRepo     interfaces.RoleRepo
	password     interfaces.PasswordHasher
	tokenService interfaces.TokenService
	tokenRepo    interfaces.TokenRepo
}

func New(
	log *slog.Logger,
	accountRepo interfaces.AccountRepo,
	roleRepo interfaces.RoleRepo,
	password interfaces.PasswordHasher,
	tokenService interfaces.TokenService,
	tokenRepo interfaces.TokenRepo,
) *Usecase {
	return &Usecase{
		log:          log,
		accountRepo:  accountRepo,
		roleRepo:     roleRepo,
		password:     password,
		tokenService: tokenService,
		tokenRepo:    tokenRepo,
	}
}
