package redistoken

import (
	"context"
	"fmt"
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (r *tokenRepo) Set(ctx context.Context, token *domain.RefreshToken) error {
	key := fmt.Sprintf("%s:%s", refreshTokenKeyPrefix, token.TokenID.String())

	err := r.db.Set(ctx, key, token.UserID.String(), time.Until(token.ExpiresAt)).Err()

	if err != nil {
		return domain.NewInternalErr(err)
	}

	return nil
}
