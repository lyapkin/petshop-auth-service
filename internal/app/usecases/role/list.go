package role

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) List(ctx context.Context) ([]domain.Role, error) {
	u.log.InfoContext(ctx, "getting list of roles started")

	roles, err := u.roleRepo.List(ctx)
	if err != nil {
		u.log.ErrorContext(ctx, "retreiving list of roles from db failed", slog.String("err", err.Error()))
		return nil, err
	}

	u.log.InfoContext(ctx, "getting list of roles succeeded")
	return roles, nil
}
