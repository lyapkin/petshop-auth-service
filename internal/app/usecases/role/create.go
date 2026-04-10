package role

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) Create(ctx context.Context, input *domain.Role) (*domain.Role, error) {
	u.log.InfoContext(ctx, "role creation starts")

	if err := input.Validate(); err != nil {
		u.log.InfoContext(ctx, "role data is not valie", slog.String("err", err.Error()))
		return nil, err
	}

	role, err := u.roleRepo.Create(ctx, input)
	if err != nil {
		u.log.InfoContext(ctx, "role db insert failed", slog.String("err", err.Error()))
		return nil, err
	}

	u.log.InfoContext(ctx, "role creation succeeded")
	return role, nil
}
