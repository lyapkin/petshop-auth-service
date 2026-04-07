package jwttoken

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) GenerateRefresh(now time.Time, account *domain.Account) (*domain.RefreshToken, error) {
	expiresAt := now.Add(s.refreshTTL)

	b := make([]byte, 32) // 32 байта = 256 бит
	_, err := rand.Read(b)
	if err != nil {
		return nil, domain.NewInternalErr(err)
	}
	token := hex.EncodeToString(b)

	return &domain.RefreshToken{
		Token:     token,
		ExpiresAt: expiresAt,
		AccountID: account.ID,
	}, nil
}
