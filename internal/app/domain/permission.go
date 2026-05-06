package domain

type Permission struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

func (r *Permission) Validate() error {
	// TODO: implement
	return nil
}
