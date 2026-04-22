package account

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (uc *Usecase) Create(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	uc.log.InfoContext(ctx, "account creation started")

	if err := input.Validate(); err != nil {
		uc.log.InfoContext(ctx, "account validation failed", slog.String("err", err.Error()))
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		uc.log.ErrorContext(ctx, "account id generation failed", slog.String("err", err.Error()))
		return nil, &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	input.ID = id
	uc.log.InfoContext(ctx, "account id generation succeeded")

	result, err := uc.accountRepo.Create(ctx, input)
	if err != nil {
		uc.log.InfoContext(ctx, "account repo storing failed", slog.String("err", err.Error()))
		return nil, err
	}

	uc.log.InfoContext(ctx, "account creation succeeded")
	return result, nil
}
