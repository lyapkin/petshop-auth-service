package pgpermission

import "database/sql"

var table string = "permission"

type permissionRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *permissionRepo {
	return &permissionRepo{
		db: db,
	}
}
