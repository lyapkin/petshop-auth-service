package account

import (
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type Usecase struct {
	log         *slog.Logger
	accountRepo interfaces.AccountRepo
}

func New(
	log *slog.Logger,
	accountRepo interfaces.AccountRepo,
) *Usecase {
	return &Usecase{
		log:         log,
		accountRepo: accountRepo,
	}
}
