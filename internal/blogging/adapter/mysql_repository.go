package adapter

import "github.com/jmoiron/sqlx"

type MySQLVoucherRepository struct {
	db *sqlx.DB
}

func NewMySQLVoucherRepository(db *sqlx.DB) *MySQLVoucherRepository {
	if db == nil {
		panic("missing db")
	}

	return &MySQLVoucherRepository{db: db}
}
