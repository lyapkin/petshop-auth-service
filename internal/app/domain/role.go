package domain

type Role struct {
	ID          int          `json:"id"`
	Slug        string       `json:"slug"`
	Name        string       `json:"name,omitempty"`
	IsBase      bool         `json:"isBase,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

func (r *Role) Validate() error {
	// TODO: implement
	return nil
}
