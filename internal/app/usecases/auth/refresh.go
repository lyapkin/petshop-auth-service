package auth

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *usecase) Refresh(ctx context.Context, refreshToken string) (*domain.Token, error) {
	u.log.InfoContext(ctx, "refresh starts")

	tokenID, err := u.tokenService.ParseRefresh(refreshToken)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to parse refresh token", slog.String("err", err.Error()))
		return nil, err
	}

	userID, err := u.tokenRepo.PopByID(ctx, tokenID)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to pop refresh token in db", slog.String("err", err.Error()))
		return nil, err
	}

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		u.log.ErrorContext(ctx, "faild to find user in db", slog.String("err", err.Error()))
		return nil, err
	}

	token, err := u.setupToken(ctx, user)

	u.log.InfoContext(ctx, "refresh finished")

	return token, nil
}
