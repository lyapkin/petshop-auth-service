package pgaccount

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *accountRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	query := `SELECT id, name, email FROM account WHERE id = $1`

	var account domain.Account
	err := r.db.QueryRowContext(ctx, query, id).Scan(&account.ID, &account.Name, &account.Email)
	if err != nil {
		return nil, buildErr(err)
	}

	return &account, nil
}
