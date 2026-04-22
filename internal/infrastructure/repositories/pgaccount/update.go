package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) Update(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	query := `
	UPDATE account SET name = $2, email = $3
	WHERE id = $1
	RETURNING id, name, email
	`

	var result domain.Account
	if err := r.db.QueryRowContext(ctx, query, input.Name, input.Email).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	// TODO: update account role

	return &result, nil
}
