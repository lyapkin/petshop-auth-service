package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) List(ctx context.Context) ([]domain.Account, error) {
	query := `
	SELECT id, name, email FROM account
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}
	defer rows.Close()

	result := make([]domain.Account, 0, 12)
	for rows.Next() {
		i := len(result)
		result := append(result, domain.Account{})
		if err := rows.Scan(result[i].ID, result[i].Name, result[i].Email); err != nil {
			return nil, postgres.BuildErr(err, table)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	return result, nil
}
