package role

import (
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type Usecase struct {
	log      *slog.Logger
	roleRepo interfaces.RoleRepo
}

func New(
	log *slog.Logger,
	roleRepo interfaces.RoleRepo,
) *Usecase {
	return &Usecase{
		log:      log,
		roleRepo: roleRepo,
	}
}
