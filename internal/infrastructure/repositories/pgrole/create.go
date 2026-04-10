package pgrole

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) Create(ctx context.Context, input *domain.Role) (*domain.Role, error) {
	query := `INSERT INTO role (slug, name, is_base) VALUES ($1, $2, $3)
						RETURNING id, slug, name, is_base`

	var role domain.Role
	if err := r.db.QueryRowContext(ctx, query, input.Slug, input.Name, input.IsBase).Scan(
		&role.ID,
		&role.Slug,
		&role.Name,
		&role.IsBase,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	// TODO: insert role_permission to connect role to its permissions

	// TODO: reset base role cache

	return &role, nil
}
