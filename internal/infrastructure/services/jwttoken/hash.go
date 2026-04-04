package jwttoken

import (
	"crypto/sha256"
	"encoding/hex"
)

func (s *service) Hash(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
