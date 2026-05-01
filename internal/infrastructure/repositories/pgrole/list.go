package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) List(ctx context.Context) ([]domain.Role, error) {
	query := `
	SELECT r.id, r.slug, r.name, r.is_base, p.id, p.slug FROM role
	JOIN role_permission rp ON r.id = rp.role_id
	JOIN permission p ON p.id = rp.permission_id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}
	defer rows.Close()

	result := make([]domain.Role, 0, 12)
	role := domain.Role{}
	permission := domain.Permission{}
	var i int
	for rows.Next() {
		if err := rows.Scan(
			&role.ID,
			&role.Slug,
			&role.Name,
			&role.IsBase,
			&permission.ID,
			&permission.Slug,
		); err != nil {
			return nil, postgres.BuildErr(err, table)
		}

		if len(result) == 0 || role.ID != result[i].ID {
			i = len(result)
			result = append(result, role)
			result[i].Permissions = make([]domain.Permission, 0)
		}

		if permission.ID != 0 {
			result[i].Permissions = append(result[i].Permissions, permission)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	return result, nil
}
