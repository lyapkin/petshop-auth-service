package auth

import (
	"context"
	"log/slog"
)

func (u *Usecase) Logout(ctx context.Context, refreshToken string) error {
	u.log.InfoContext(ctx, "logout starts")

	hash := u.tokenService.Hash(refreshToken)

	_, err := u.tokenRepo.Pop(ctx, hash)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to pop refresh token in db", slog.String("err", err.Error()))
		return err
	}

	return nil
}
