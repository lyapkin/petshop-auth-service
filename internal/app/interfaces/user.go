package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type AccountRepo interface {
	Create(context.Context, *domain.Account) (*domain.Account, error)
	GetByLogin(context.Context, string) (*domain.Account, error)
	GetByID(context.Context, uuid.UUID) (*domain.Account, error)
}
