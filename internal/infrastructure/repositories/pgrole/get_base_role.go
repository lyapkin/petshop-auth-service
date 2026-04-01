package pgrole

import (
	"context"
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *roleRepo) GetBaseRole(ctx context.Context) (*domain.Role, error) {
	// check cache
	r.mu.RLock()
	if r.cachedRole != nil && time.Now().Before(r.expiresAt) {
		role := r.cachedRole
		r.mu.RUnlock()
		return role, nil
	}
	r.mu.RUnlock()

	// take lock to cache value
	r.mu.Lock()
	defer r.mu.Unlock()

	// check if another goroutine has cached value while this one was waiting for the lock
	if r.cachedRole != nil && time.Now().Before(r.expiresAt) {
		return r.cachedRole, nil
	}

	// go to db
	query := `SELECT id FROM role WHERE is_base = true`
	var role domain.Role
	err := r.db.QueryRowContext(ctx, query).Scan(&role.ID)
	if err != nil {
		return nil, domain.NewInternalErr(err)
	}

	// cache the value
	r.cachedRole = &role
	r.expiresAt = time.Now().Add(r.ttl)

	return r.cachedRole, nil
}
