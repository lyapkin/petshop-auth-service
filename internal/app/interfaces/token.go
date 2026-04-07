package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type TokenService interface {
	GeneratePair(account *domain.Account) (*domain.Token, error)
	Hash(token string) string
}

type TokenRepo interface {
	Set(ctx context.Context, hash string, token *domain.RefreshToken) error
	Pop(context.Context, string) (accountID uuid.UUID, err error)
}
