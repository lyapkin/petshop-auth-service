package pgaccount

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	query := `
	SELECT a.id, a.name, a.email, r.id, r.slug FROM account a
	LEFT JOIN account_role ar ON a.id = ar.account_id
	LEFT JOIN role r ON ar.role_id = r.id
	WHERE a.id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, buildErr(err)
	}
	defer rows.Close()

	account := domain.Account{
		Roles: make([]domain.Role, 0),
	}

	var rID sql.NullInt64
	var rSlug sql.NullString

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

		if rID.Valid {
			account.Roles = append(account.Roles, domain.Role{
				ID:   int(rID.Int64),
				Slug: rSlug.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	return &account, nil
}
