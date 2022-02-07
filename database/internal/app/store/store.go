package store

import (
	"database/sql"
	"log"
	"os"

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

func (s *Store) Customer() *CustomerProvider {
	return &CustomerProvider{
		store: s,
		l:     log.New(os.Stdout, "customer:", log.LstdFlags|log.LUTC),
	}
}

func (s *Store) Close() error {
	return s.db.Close()
}
