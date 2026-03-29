package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Role     []Role    `json:"roles"`
}

func (u *User) Validate() error {
	// TODO: implement
	return nil
}
