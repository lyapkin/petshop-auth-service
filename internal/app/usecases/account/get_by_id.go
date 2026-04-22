package account

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (uc *Usecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	uc.log.InfoContext(ctx, "account retreiving by id started")

	result, err := uc.accountRepo.GetByID(ctx, id)
	if err != nil {
		uc.log.WarnContext(ctx, "account db retreiving by id failed", slog.String("err", err.Error()))
		return nil, err
	}

	uc.log.InfoContext(ctx, "account retreiving by id succeeded")
	return result, nil
}
