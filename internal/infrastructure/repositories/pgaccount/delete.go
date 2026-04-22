package pgaccount

import (
	"context"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *accountRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM account WHERE id = $1"

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return postgres.BuildErr(err, table)
	}

	return nil
}
