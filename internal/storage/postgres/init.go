package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/lyapkin/shop/auth/config"
)

func New(cfg config.DB) (*sql.DB, error) {
	db, err := sql.Open(string(cfg.Driver), cfg.URL())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
