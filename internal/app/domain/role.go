package domain

type Role struct {
	ID          int
	Slug        string
	Name        string
	IsBase      bool
	Permissions []Permission
}
