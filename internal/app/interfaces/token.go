package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type TokenService interface {
	GeneratePair(user *domain.User, tokenID uuid.UUID) (*domain.Token, error)
	ParseRefresh(token string) (tokenID uuid.UUID, err error)
}

type TokenRepo interface {
	Set(ctx context.Context, tokenID uuid.UUID, userID uuid.UUID) error
	PopByID(context.Context, uuid.UUID) (userID uuid.UUID, err error)
}
