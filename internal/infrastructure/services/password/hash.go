package password

import "github.com/alexedwards/argon2id"

func (s *service) GenerateHash(pass string) (string, error) {
	return argon2id.CreateHash(pass, &argon2id.Params{
		Memory:      memory,
		Iterations:  iterations,
		Parallelism: threads,
		SaltLength:  saltLength,
		KeyLength:   keyLength,
	})
}
