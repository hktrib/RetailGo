package db

import "database/sql"

type Store struct {
	*Queries
	DB *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{
		DB:      db,
		Queries: New(db),
	}
}
