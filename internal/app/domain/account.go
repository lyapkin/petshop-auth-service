package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Roles    []Role    `json:"roles"`
}

func (u *Account) Validate() error {
	// TODO: implement
	return nil
}
