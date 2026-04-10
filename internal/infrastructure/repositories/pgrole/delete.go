package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) Delete(ctx context.Context, id int) error {
	query := `
	DELETE FROM role WHERE id = $1
	`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return postgres.BuildErr(err, table)
	}

	return nil
}
