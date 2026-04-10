package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func BuildErr(err error, table string) error {
	code := domain.ErrInternal
	message := domain.InternalErrorMessage

	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code.Name() == "unique_violation" {

		switch pgErr.Column {
		case "id":
			message = "Internal server error. Try again."
		default:
			code = domain.ErrDuplicate
			message = fmt.Sprintf("%s with the field already exists: %s", table, pgErr.Column)
		}
	}

	if err == sql.ErrNoRows {
		code = domain.ErrNotFound
		message = fmt.Sprintf("%s does not exist", table)
	}

	return &domain.AppError{
		Code:     code,
		Message:  message,
		Internal: err,
	}
}
