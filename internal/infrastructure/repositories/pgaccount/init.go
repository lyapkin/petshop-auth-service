package pgaccount

import "database/sql"

var table string = "account"

type accountRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}
