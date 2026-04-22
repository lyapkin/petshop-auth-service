package pgaccount

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) Update(ctx context.Context, input *domain.Account) (*domain.Account, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, postgres.BuildErr(err, "")
	}
	defer tx.Rollback()

	query := `
	UPDATE account SET name = $2, email = $3
	WHERE id = $1
	RETURNING id, name, email
	`
	var result domain.Account
	if err := tx.QueryRowContext(ctx, query, input.ID, input.Name, input.Email).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	query = `
	DELETE FROM account_role WHERE account_id = $1
	`
	_, err = tx.ExecContext(ctx, query, input.ID)
	if err != nil {
		return nil, postgres.BuildErr(err, "account_role")
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

	result.Roles = input.Roles

	return &result, nil
}
