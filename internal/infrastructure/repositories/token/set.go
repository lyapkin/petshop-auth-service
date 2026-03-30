package token

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *tokenRepo) Set(ctx context.Context, tokenID uuid.UUID, userID uuid.UUID, expiration time.Time) error {
	key := fmt.Sprintf("%s:%s", refreshTokenKeyPrefix, tokenID.String())

	err := r.db.Set(ctx, key, userID.String(), time.Until(expiration)).Err()

	if err != nil {
		return domain.NewInternalErr(err)
	}

	return nil
}
