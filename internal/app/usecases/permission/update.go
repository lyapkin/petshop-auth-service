package permission

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) Update(ctx context.Context, input *domain.Permission) (*domain.Permission, error) {
	u.log.InfoContext(ctx, "permission update started")

	if err := input.Validate(); err != nil {
		u.log.InfoContext(ctx, "permission data is not valid", slog.String("err", err.Error()))
		return nil, err
	}

	permission, err := u.permissionRepo.Update(ctx, input)
	if err != nil {
		u.log.InfoContext(ctx, "permission db update failed", slog.String("err", err.Error()))
		return nil, err
	}

	u.log.InfoContext(ctx, "permission updated succeeded")
	return permission, nil
}
