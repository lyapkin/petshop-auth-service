package pgpermission

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *permissionRepo) Delete(ctx context.Context, id int) error {
	query := `
	DELETE FROM permission WHERE id = $1
	`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return postgres.BuildErr(err, table)
	}

	return nil
}
