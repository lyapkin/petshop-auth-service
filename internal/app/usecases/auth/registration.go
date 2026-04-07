package auth

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (u *Usecase) Register(ctx context.Context, account *domain.Account) (err error) {
	u.log.InfoContext(ctx, "account resigstration started")

	if err = account.Validate(); err != nil {
		u.log.InfoContext(ctx, "account data is not valie", slog.String("err", err.Error()))
		return err
	}
	u.log.InfoContext(ctx, "account data validation succeeded")

	// generate account id
	account.ID, err = uuid.NewRandom()
	if err != nil {
		u.log.ErrorContext(ctx, "account id generation failed", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	u.log.InfoContext(ctx, "account id generation succeeded")

	// hash account password
	account.Password, err = u.password.Hash(account.Password)
	if err != nil {
		u.log.ErrorContext(ctx, "account password hashing failed", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	u.log.InfoContext(ctx, "account password hashing succeeded")

	// setting base role
	role, err := u.roleRepo.GetBaseRole(ctx)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to rerieve base role", slog.String("err", err.Error()))
		return &domain.AppError{
			Code:     domain.ErrInternal,
			Message:  domain.InternalErrorMessage,
			Internal: err,
		}
	}
	u.log.InfoContext(ctx, "retrieving base role succeeded")
	account.Roles = append(account.Roles, *role)

	// storing account to db
	_, err = u.accountRepo.Create(ctx, account)
	if err != nil {
		u.log.InfoContext(ctx, "account db insert failed", slog.String("err", err.Error()))
		return err
	}

	u.log.InfoContext(ctx, "account registration succeeded")
	return nil
}
