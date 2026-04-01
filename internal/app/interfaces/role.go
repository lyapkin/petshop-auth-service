package interfaces

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type RoleRepo interface {
	GetBaseRole(context.Context) (*domain.Role, error)
}
