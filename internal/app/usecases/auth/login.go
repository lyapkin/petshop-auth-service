package auth

import (
	"context"
	"errors"
	"log/slog"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/app/dto"
)

func (u *Usecase) Login(ctx context.Context, input *dto.LoginInput) (*domain.Token, error) {
	u.log.InfoContext(ctx, "login starts")

	user, err := u.userRepo.GetByLogin(ctx, input.Login)
	if err != nil {
		u.log.InfoContext(ctx, "failed to retreive user from db", slog.String("err", err.Error()))
		return nil, err
	}

	match, err := u.password.Compare(input.Password, user.Password)
	if err != nil {
		u.log.ErrorContext(ctx, "login password checking failed", slog.String("err", err.Error()))
		return nil, err
	}

	if !match {
		u.log.InfoContext(ctx, "login password not matched")
		return nil, &domain.AppError{
			Code:     domain.ErrNotAuthenticated,
			Message:  "login or password not exist",
			Internal: errors.New("password match failed"),
		}
	}

	token, err := u.setupToken(ctx, user)

	u.log.InfoContext(ctx, "login finished")

	return token, nil
}
