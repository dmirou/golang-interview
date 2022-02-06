package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New(databaseURL string) (*Store, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) User() *UserProvider {
	return &UserProvider{
		store: s,
	}
}

func (s *Store) Close() error {
	return s.db.Close()
}
