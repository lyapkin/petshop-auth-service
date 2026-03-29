package interfaces

type PasswordHasher interface {
	Hash(password string) (hash string, err error)
	Compare(password, hash string) (match bool, err error)
}
