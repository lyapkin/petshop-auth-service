package auth

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) Refresh(ctx context.Context, refreshToken string) (*domain.Token, error) {
	u.log.InfoContext(ctx, "refresh starts")

	hash := u.tokenService.Hash(refreshToken)

	accountID, err := u.tokenRepo.Pop(ctx, hash)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to pop refresh token in db", slog.String("err", err.Error()))
		return nil, err
	}

	account, err := u.accountRepo.GetByID(ctx, accountID)
	if err != nil {
		u.log.ErrorContext(ctx, "faild to find account in db", slog.String("err", err.Error()))
		return nil, err
	}

	token, err := u.setupToken(ctx, account)

	u.log.InfoContext(ctx, "refresh finished")

	return token, nil
}
