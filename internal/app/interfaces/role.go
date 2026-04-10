package interfaces

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type RoleRepo interface {
	GetBaseRole(context.Context) (*domain.Role, error)
	Create(context.Context, *domain.Role) (*domain.Role, error)
	Update(context.Context, *domain.Role) (*domain.Role, error)
	Delete(context.Context, int) error
	List(context.Context) ([]domain.Role, error)
}
