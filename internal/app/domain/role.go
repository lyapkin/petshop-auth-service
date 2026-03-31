package domain

type Role struct {
	ID          int
	Slug        string
	Name        string
	Permissions []Permission
}
