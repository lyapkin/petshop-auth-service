package domain

type Role struct {
	ID          int          `json:"id"`
	Slug        string       `json:"slug"`
	Name        string       `json:"name"`
	IsBase      bool         `json:"isBase"`
	Permissions []Permission `json:"permissions"`
}

func (r *Role) Validate() error {
	// TODO: implement
	return nil
}
