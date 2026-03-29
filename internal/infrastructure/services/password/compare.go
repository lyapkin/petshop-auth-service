package password

import "github.com/alexedwards/argon2id"

func ComparePasswordAndHash(password, hash string) (match bool, err error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
