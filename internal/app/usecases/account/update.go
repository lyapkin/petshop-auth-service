package account

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (uc *Usecase) Update(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	uc.log.InfoContext(ctx, "account update started")

	if err := input.Validate(); err != nil {
		uc.log.InfoContext(ctx, "account data validation failed", slog.String("err", err.Error()))
		return nil, err
	}

	result, err := uc.accountRepo.Update(ctx, input)
	if err != nil {
		uc.log.InfoContext(ctx, "account db update failed", slog.String("err", err.Error()))
		return nil, err
	}

	uc.log.InfoContext(ctx, "account update succeeded")
	return result, nil
}
