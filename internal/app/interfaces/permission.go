package interfaces

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type PermissionRepo interface {
	Create(context.Context, *domain.Permission) (*domain.Permission, error)
	Update(context.Context, *domain.Permission) (*domain.Permission, error)
	Delete(context.Context, int) error
	List(context.Context) ([]domain.Permission, error)
}
