package pgaccount

import "database/sql"

type accountRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}
