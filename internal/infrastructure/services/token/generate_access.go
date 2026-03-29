package token

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

func (s *service) GenerateAccess(now time.Time, user *domain.User) (string, error) {
	accessClaims := accessTokenClaims{
		Name:   user.Name,
		Scopes: getUserUniquePermissions(user),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims).SignedString(s.accessSecret)
	if err != nil {
		return "", domain.NewInternalErr(err)
	}

	return accessToken, nil
}

func getUserUniquePermissions(u *domain.User) []domain.Permission {
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
