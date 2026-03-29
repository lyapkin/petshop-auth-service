package auth

import (
	"context"
	"log/slog"
)

func (u *usecase) Logout(ctx context.Context, refreshToken string) error {
	u.log.InfoContext(ctx, "logout starts")

	tokenID, err := u.tokenService.ParseRefresh(refreshToken)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to parse refresh token", slog.String("err", err.Error()))
		return err
	}

	_, err = u.tokenRepo.PopByID(ctx, tokenID)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to pop refresh token in db", slog.String("err", err.Error()))
		return err
	}

	return nil
}
