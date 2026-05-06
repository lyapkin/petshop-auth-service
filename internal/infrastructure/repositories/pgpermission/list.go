package pgpermission

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *permissionRepo) List(ctx context.Context) ([]domain.Permission, error) {
	query := `
	SELECT id, slug FROM permission
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	result := make([]domain.Permission, 0, 12)
	for rows.Next() {
		i := len(result)
		result = append(result, domain.Permission{})
		if err := rows.Scan(&result[i].ID, &result[i].Slug); err != nil {
			return nil, postgres.BuildErr(err, table)
		}
	}

	return result, nil
}
