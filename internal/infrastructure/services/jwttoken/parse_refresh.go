package jwttoken

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) ParseRefresh(refreshToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong algorithm: %v", token.Header["alg"])
		}
		return s.refreshSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, &domain.AppError{
				Code:     domain.ErrNotAuthenticated,
				Message:  "token expired",
				Internal: err,
			}
		}

		if errors.Is(err, jwt.ErrTokenSignatureInvalid) ||
			errors.Is(err, jwt.ErrTokenMalformed) ||
			errors.Is(err, jwt.ErrTokenUnverifiable) {
			return uuid.Nil, &domain.AppError{
				Code:     domain.ErrNotValid,
				Message:  "invalid token structure or signature",
				Internal: err,
			}
		}

		return uuid.Nil, domain.NewInternalErr(err)
	}

	claims, ok := token.Claims.(*refreshTokenClaims)
	if !ok {
		return uuid.Nil, domain.NewInternalErr(errors.New("cannot assert type of refresh token claims"))
	}

	return claims.TokenID, nil
}
