package auth

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *usecase) Register(ctx context.Context, user *domain.User) (err error) {
	u.log.InfoContext(ctx, "user resigstration started")

	if err = user.Validate(); err != nil {
		u.log.InfoContext(ctx, "user data is not valie", slog.String("err", err.Error()))
		return err
	}
	u.log.InfoContext(ctx, "user data validation succeeded")

	// generate user id
	user.ID, err = uuid.NewRandom()
	if err != nil {
		u.log.ErrorContext(ctx, "user id generation failed", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	u.log.InfoContext(ctx, "user id generation succeeded")

	// hash user password
	user.Password, err = u.password.Hash(user.Password)
	if err != nil {
		u.log.ErrorContext(ctx, "user password hashing failed", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	u.log.InfoContext(ctx, "user password hashing succeeded")

	// setting base role
	role, err := u.roleRepo.FindBaseRole()
	if err != nil {
		u.log.ErrorContext(ctx, "finding base role failed", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	user.Role = append(user.Role, *role)
	u.log.InfoContext(ctx, "user role setting succeeded")

	// storing user to db
	_, err = u.userRepo.Create(ctx, user)
	if err != nil {
		u.log.InfoContext(ctx, "user db insert failed", slog.String("err", err.Error()))
		return err
	}

	u.log.InfoContext(ctx, "user registration succeeded")
	return nil
}
