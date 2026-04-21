package permission

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) List(ctx context.Context) ([]domain.Permission, error) {
	u.log.InfoContext(ctx, "getting list of permissions started")

	result, err := u.permissionRepo.List(ctx)
	if err != nil {
		u.log.ErrorContext(ctx, "retreiving list of permissions from db failed", slog.String("err", err.Error()))
		return nil, err
	}

	u.log.InfoContext(ctx, "getting list of permissions succeeded")
	return result, nil
}
