package auth

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/app/interfaces"
)

type usecase struct {
	log          *slog.Logger
	userRepo     interfaces.UserRepo
	roleRepo     interfaces.RoleRepo
	password     interfaces.PasswordHasher
	tokenService interfaces.TokenService
	tokenRepo    interfaces.TokenRepo
}

func New(
	log *slog.Logger,
	userRepo interfaces.UserRepo,
	roleRepo interfaces.RoleRepo,
	password interfaces.PasswordHasher,
	tokenService interfaces.TokenService,
	tokenRepo interfaces.TokenRepo,
) *usecase {
	return &usecase{
		log:          log,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		password:     password,
		tokenService: tokenService,
		tokenRepo:    tokenRepo,
	}
}

func (u *usecase) setupToken(ctx context.Context, user *domain.User) (*domain.Token, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		u.log.ErrorContext(ctx, "token id generation failed", slog.String("err", err.Error()))
		return nil, domain.NewInternalErr(err)
	}
	u.log.InfoContext(ctx, "refresh tokenID generation succeeded")

	token, err := u.tokenService.GeneratePair(user, tokenID)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to genereate tokens", slog.String("err", err.Error()))
		return nil, err
	}
	u.log.InfoContext(ctx, "tokens generation succeeded")

	err = u.tokenRepo.Set(ctx, tokenID, user.ID)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to store refresh token", slog.String("err", err.Error()))
		return nil, err
	}
	u.log.InfoContext(ctx, "store refresh token succeeded")

	return token, nil
}
