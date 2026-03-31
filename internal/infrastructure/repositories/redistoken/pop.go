package redistoken

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/redis/go-redis/v9"
)

func (r *tokenRepo) PopByID(ctx context.Context, tokenID uuid.UUID) (userID uuid.UUID, err error) {
	key := fmt.Sprintf("%s:%s", refreshTokenKeyPrefix, tokenID.String())

	result, err := r.db.GetDel(ctx, key).Result()
	if err == redis.Nil {
		return uuid.Nil, &domain.AppError{
			Code:     domain.ErrNotAuthenticated,
			Message:  "not authenticated",
			Internal: err,
		}
	} else if err != nil {
		return uuid.Nil, domain.NewInternalErr(err)
	}

	userID, err = uuid.Parse(result)
	if err != nil {
		return uuid.Nil, domain.NewInternalErr(err)
	}

	return userID, nil
}
