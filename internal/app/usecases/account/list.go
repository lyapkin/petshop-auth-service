package account

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (uc *Usecase) List(ctx context.Context) ([]domain.Account, error) {
	uc.log.InfoContext(ctx, "account list retreiving started")

	result, err := uc.accountRepo.List(ctx)
	if err != nil {
		uc.log.ErrorContext(ctx, "account db list retreiving failed", slog.String("err", err.Error()))
		return nil, err
	}

	uc.log.InfoContext(ctx, "account list retreiving succeeded")
	return result, nil
}
