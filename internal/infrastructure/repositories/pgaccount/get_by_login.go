package pgaccount

import (
	"context"
	"database/sql"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) GetByLogin(ctx context.Context, login string) (*domain.Account, error) {
	query := `SELECT id, name, email, password FROM account WHERE email = $1`

	var account domain.Account
	err := r.db.QueryRowContext(ctx, query, login).Scan(&account.ID, &account.Name, &account.Email, &account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &domain.AppError{
				Code:     domain.ErrNotAuthenticated,
				Message:  "login or password not exist",
				Internal: err,
			}
		}
		return nil, postgres.BuildErr(err, table)
	}

	return &account, nil
}
