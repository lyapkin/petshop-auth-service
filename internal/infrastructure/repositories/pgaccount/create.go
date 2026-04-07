package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *accountRepo) Create(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	query := `INSERT INTO account (id, name, email, password)
	VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, input.ID, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, buildErr(err)
	}

	return input, nil
}
