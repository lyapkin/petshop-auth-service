package pgrole

import (
	"database/sql"
	"sync"
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

var table string = "Role"

type roleRepo struct {
	db         *sql.DB
	mu         sync.RWMutex
	cachedRole *domain.Role
	expiresAt  time.Time
	ttl        time.Duration
}

func New(db *sql.DB, ttl time.Duration) *roleRepo {
	return &roleRepo{
		db:  db,
		ttl: ttl,
	}
}
