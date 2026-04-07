package jwttoken

import (
	"time"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func (s *service) GeneratePair(account *domain.Account) (*domain.Token, error) {
	now := time.Now()

	accessToken, err := s.GenerateAccess(now, account)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefresh(now, account)
	if err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}
