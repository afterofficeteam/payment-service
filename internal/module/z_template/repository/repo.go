package repository

import "database/sql"

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

func (r *store) CreatePayment(string) (string, error) {
	return "", nil
}
