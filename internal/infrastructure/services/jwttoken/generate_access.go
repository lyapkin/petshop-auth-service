package jwttoken

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lyapkin/shop/auth/internal/app/domain"
)

type accessTokenClaims struct {
	Name   string              `json:"name"`
	Scopes []domain.Permission `json:"scopes"`
	jwt.RegisteredClaims
}

func (s *service) GenerateAccess(now time.Time, account *domain.Account) (*domain.AccessToken, error) {
	accessClaims := accessTokenClaims{
		Name:   account.Name,
		Scopes: getAccountUniquePermissions(account),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   account.ID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims).SignedString(s.accessSecret)
	if err != nil {
		return nil, domain.NewInternalErr(err)
	}

	return &domain.AccessToken{
		Token: accessToken,
	}, nil
}

func getAccountUniquePermissions(u *domain.Account) []domain.Permission {
	seen := make(map[int]struct{}, 12)

	permissions := make([]domain.Permission, 0, 12)
	for _, role := range u.Roles {
		for _, permission := range role.Permissions {
			if _, ok := seen[permission.ID]; !ok {
				permissions = append(permissions, permission)
			}
		}
	}

	return permissions
}
