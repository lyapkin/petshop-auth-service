package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type UserRepo interface {
	Create(context.Context, *domain.User) (*domain.User, error)
	GetByLogin(context.Context, string) (*domain.User, error)
	GetByID(context.Context, uuid.UUID) (*domain.User, error)
}
