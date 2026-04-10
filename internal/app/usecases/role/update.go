package role

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) Update(ctx context.Context, input *domain.Role) (*domain.Role, error) {
	u.log.InfoContext(ctx, "role update started")

	if err := input.Validate(); err != nil {
		u.log.InfoContext(ctx, "role data is not valie", slog.String("err", err.Error()))
		return nil, err
	}

	role, err := u.roleRepo.Update(ctx, input)
	if err != nil {
		u.log.ErrorContext(ctx, "role db update failed", slog.String("err", err.Error()))
		return nil, err
	}

	u.log.InfoContext(ctx, "role update succeeded")
	return role, nil
}
