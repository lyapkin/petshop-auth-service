package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *accountRepo) GetByLogin(ctx context.Context, login string) (*domain.Account, error) {
	query := `SELECT id, name, email FROM account WHERE email = $1`

	var account domain.Account
	err := r.db.QueryRowContext(ctx, query, login).Scan(&account.ID, &account.Name, &account.Email)
	if err != nil {
		return nil, buildErr(err)
	}

	return &account, nil
}
