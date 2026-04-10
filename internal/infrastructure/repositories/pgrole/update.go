package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) Update(ctx context.Context, input *domain.Role) (*domain.Role, error) {
	query := `
	UPDATE role SET slug = $2, name = $3 is_base = $4
	WHERE id = $1
	RETURNING id, slug, name, is_base
	`

	var role domain.Role
	if err := r.db.QueryRowContext(ctx, query, input.ID, input.Slug, input.Name, input.IsBase).Scan(
		&role.ID,
		&role.Slug,
		&role.Name,
		&role.IsBase,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	// TODO: update role_permission

	// TODO: reset base role cache

	return &role, nil
}
