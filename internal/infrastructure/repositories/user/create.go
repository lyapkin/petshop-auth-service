package user

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *userRepo) Create(ctx context.Context, input *domain.User) (*domain.User, error) {
	query := `INSERT INTO user (id, name, email, password)
	VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, input.ID, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, buildErr(err)
	}

	return input, nil
}
