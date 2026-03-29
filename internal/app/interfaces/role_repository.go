package interfaces

import "github.com/lyapkin/shop/auth/internal/app/domain"

type RoleRepo interface {
	FindBaseRole() (*domain.Role, error)
}
