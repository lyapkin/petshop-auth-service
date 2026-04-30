package pgrole

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
)

func (r *roleRepo) Delete(ctx context.Context, id int) error {
	query := `
	WITH deleted AS (
		DELETE FROM role WHERE id = $1 AND is_base = false RETURNING id
	)
	SELECT is_base FROM role WHERE id = $1 AND NOT EXISTS (SELECT 1 FROM deleted)
	`
	var isBase bool
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&isBase); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return postgres.BuildErr(err, table)
	}

	if isBase {
		return &domain.AppError{
			Code:     domain.ErrBuisnessRuleViolation,
			Message:  "can't delete base role",
			Internal: fmt.Errorf("role is base"),
		}
	}

	return nil
}
