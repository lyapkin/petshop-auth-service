package domain

type Role struct {
	ID          int
	Slug        string
	Name        string
	IsBase      bool
	Permissions []Permission
}

func (r *Role) Validate() error {
	// TODO: implement
	return nil
}
