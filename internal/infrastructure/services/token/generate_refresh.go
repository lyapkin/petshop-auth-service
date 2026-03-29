package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) GenerateRefresh(now time.Time, user *domain.User, tokenID uuid.UUID) (string, error) {
	refreshClaims := refreshTokenClaims{
		TokenID: tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(s.refreshSecret)
	if err != nil {
		return "", domain.NewInternalErr(err)
	}

	return refreshToken, nil
}
