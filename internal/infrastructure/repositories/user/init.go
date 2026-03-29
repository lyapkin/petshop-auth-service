package user

import "database/sql"

type userRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}
