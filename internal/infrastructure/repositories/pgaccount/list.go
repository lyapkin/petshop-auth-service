package pgaccount

import (
	"context"
	"database/sql"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) List(ctx context.Context) ([]domain.Account, error) {
	query := `
	SELECT a.id, a.name, a.email, r.id, r.slug FROM account a
	LEFT JOIN account_role ar ON a.id = ar.account_id
	LEFT JOIN role r ON ar.role_id = r.id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}
	defer rows.Close()

	result := make([]domain.Account, 0, 12)
	account := domain.Account{}

	var rID sql.NullInt64
	var rSlug sql.NullString

	var i int
	for rows.Next() {
		if err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Email,
			&rID,
			&rSlug,
		); err != nil {
			return nil, postgres.BuildErr(err, table)
		}

		if len(result) == 0 || account.ID != result[i].ID {
			i = len(result)
			result = append(result, account)
			result[i].Roles = make([]domain.Role, 0)
		}

		if rID.Valid {
			result[i].Roles = append(result[i].Roles, domain.Role{
				ID:   int(rID.Int64),
				Slug: rSlug.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	return result, nil
}
