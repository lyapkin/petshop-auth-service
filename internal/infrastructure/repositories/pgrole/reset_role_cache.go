package pgrole

import (
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *roleRepo) resetRoleCache(role *domain.Role) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cachedRole = role
	r.expiresAt = time.Now().Add(r.ttl)
}
