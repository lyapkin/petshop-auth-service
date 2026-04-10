package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) List(ctx context.Context) ([]domain.Role, error) {
	query := `
	SELECT id, slug, name, is_base FROM role
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	result := make([]domain.Role, 0, 12)
	for rows.Next() {
		i := len(result)
		result := append(result, domain.Role{})

		if err := rows.Scan(
			&result[i].ID,
			&result[i].Slug,
			&result[i].Name,
			&result[i].IsBase,
		); err != nil {
			return nil, postgres.BuildErr(err, table)
		}
	}

	// TODO: join with Permission

	return result, nil
}
