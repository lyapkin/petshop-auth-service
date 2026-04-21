package pgpermission

import (
	"context"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *permissionRepo) Create(ctx context.Context, input *domain.Permission) (*domain.Permission, error) {
	query := `
	INSERT INTO permission (slug) VALUES ($1)
	RETURNING id, slug
	`

	var result domain.Permission
	if err := r.db.QueryRowContext(ctx, query, input.Slug).Scan(
		&result.ID,
		&result.Slug,
	); err != nil {
		return nil, postgres.BuildErr(err, table)
	}

	return &result, nil
}
