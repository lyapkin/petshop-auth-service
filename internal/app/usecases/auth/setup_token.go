package auth

import (
	"context"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) setupToken(ctx context.Context, user *domain.User) (*domain.Token, error) {
	token, err := u.tokenService.GeneratePair(user)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to genereate tokens", slog.String("err", err.Error()))
		return nil, err
	}
	u.log.InfoContext(ctx, "tokens generation succeeded")

	hash := u.tokenService.Hash(token.RefreshToken.Token)

	err = u.tokenRepo.Set(ctx, hash, &token.RefreshToken)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to store refresh token", slog.String("err", err.Error()))
		return nil, err
	}
	u.log.InfoContext(ctx, "store refresh token succeeded")

	return token, nil
}
