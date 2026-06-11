package jwttoken

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) ValidateToken(token string) (*domain.AccessTokenClaims, error) {
	claims := &accessTokenClaims{}
	result, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, &domain.AppError{
					Code:     domain.ErrNotValid,
					Message:  "token is not valid",
					Internal: errors.New("jwt token signing method is not valid"),
				}
			}

			return s.accessPublic, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, &domain.AppError{
				Code:     domain.ErrNotValid,
				Message:  "token expired",
				Internal: err,
			}
		}
	}

	if !result.Valid {
		return nil, &domain.AppError{
			Code:     domain.ErrNotValid,
			Message:  "token not valid",
			Internal: errors.New("invalid token signature"),
		}
	}

	return &claims.AccessTokenClaims, nil
}
