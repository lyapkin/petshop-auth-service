package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type accessTokenClaims struct {
	domain.AccessTokenClaims
	jwt.RegisteredClaims
}
