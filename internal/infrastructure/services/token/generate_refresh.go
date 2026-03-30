package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) GenerateRefresh(now time.Time, user *domain.User, tokenID uuid.UUID) (*domain.RefreshToken, error) {
	expiresAt := now.Add(s.refreshTTL)
	refreshClaims := refreshTokenClaims{
		TokenID: tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(s.refreshSecret)
	if err != nil {
		return nil, domain.NewInternalErr(err)
	}

	return &domain.RefreshToken{
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	}, nil
}
