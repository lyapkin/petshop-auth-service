package pguser

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *userRepo) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	query := `SELECT id, name, email FROM user WHERE email = $1`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, buildErr(err)
	}

	return &user, nil
}
