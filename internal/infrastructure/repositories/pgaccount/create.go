package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) Create(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, postgres.BuildErr(err, "")
	}
	defer tx.Rollback()

	query := `INSERT INTO account (id, name, email, password)
	VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, query, input.ID, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	query = `
	INSERT INTO account_role (account_id, role_id)
	VALUES ($1, $2)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, "account_role")
	}
	defer stmt.Close()
	for _, role := range input.Roles {
		_, err := stmt.ExecContext(ctx, input.ID, role.ID)
		if err != nil {
			return nil, postgres.BuildErr(err, "account_role")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, postgres.BuildErr(err, "")
	}

	return input, nil
}
