package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) Update(ctx context.Context, input *domain.Role) (*domain.Role, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, postgres.BuildErr(err, "")
	}
	defer tx.Rollback()

	query := `
	UPDATE role SET slug = $2, name = $3, is_base = $4
	WHERE id = $1
	RETURNING id, slug, name, is_base
	`
	var role domain.Role
	if err := tx.QueryRowContext(ctx, query, input.ID, input.Slug, input.Name, input.IsBase).Scan(
		&role.ID,
		&role.Slug,
		&role.Name,
		&role.IsBase,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	// TODO: update role_permission
	query = `
	DELETE FROM role_permission WHERE role_id = $1
	`
	if _, err := tx.ExecContext(ctx, query, input.ID); err != nil {
		return nil, postgres.BuildErr(err, joinTable)
	}

	query = `
	INSERT INTO role_permission (role_id, permission_id)
	VALUES ($1, $2)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, postgres.BuildErr(err, joinTable)
	}
	defer stmt.Close()
	for _, permission := range input.Permissions {
		if _, err := stmt.ExecContext(ctx, input.ID, permission.ID); err != nil {
			return nil, postgres.BuildErr(err, joinTable)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, postgres.BuildErr(err, "")
	}

	if role.IsBase {
		// reset base role cache
		r.resetRoleCache(&role)
	}

	return &role, nil
}
