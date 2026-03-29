package domain

import "encoding/json"

type Permission struct {
	ID   int
	Slug string
}

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Slug)
}
