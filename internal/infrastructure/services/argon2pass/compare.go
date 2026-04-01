package argon2pass

import "github.com/alexedwards/argon2id"

func (s *service) Compare(password, hash string) (match bool, err error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
