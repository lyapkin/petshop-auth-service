package permission

import (
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type Usecase struct {
	log            *slog.Logger
	permissionRepo interfaces.PermissionRepo
}

func New(log *slog.Logger, permissionRepo interfaces.PermissionRepo) *Usecase {
	return &Usecase{
		log:            log,
		permissionRepo: permissionRepo,
	}
}
