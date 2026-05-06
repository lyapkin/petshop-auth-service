package role

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (uc *Usecase) GetByID(ctx context.Context, id int) (*domain.Role, error) {
	slog.InfoContext(ctx, "getting role by id started")

	result, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to retreive role by id", slog.String("err", err.Error()))
		return nil, err
	}

	slog.InfoContext(ctx, "getting role by id succeeded")

	return result, nil
}
