package pgrole

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) GetByID(ctx context.Context, id int) (*domain.Role, error) {
	query := `
	SELECT r.id, r.slug, r.name, r.is_base, p.id, p.slug FROM role r
	LEFT JOIN role_permission rp ON r.id = rp.role_id
	LEFT JOIN permission p ON rp.permission_id = p.id
	WHERE r.id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, postgres.BuildErr(err, table)
	}
	defer rows.Close()

	role := domain.Role{
		Permissions: make([]domain.Permission, 0),
	}

	var pID sql.NullInt64
	var pSlug sql.NullString

	for rows.Next() {
		if err := rows.Scan(
			&role.ID,
			&role.Slug,
			&role.Name,
			&role.IsBase,
			&pID,
			&pSlug,
		); err != nil {
			return nil, postgres.BuildErr(err, table)
		}

		if pID.Valid {
			role.Permissions = append(role.Permissions, domain.Permission{
				ID:   int(pID.Int64),
				Slug: pSlug.String,
			})
		}
	}

	if role.ID == 0 {
		return nil, &domain.AppError{
			Code:     domain.ErrNotFound,
			Message:  "role does not exists",
			Internal: fmt.Errorf("no role with id: %d", id),
		}
	}

	return &role, nil
}
