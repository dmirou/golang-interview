package main

import (
	"log"
	"os"

	"github.com/dmirou/learngo/database/internal/app/store"
)

func main() {
	connStr := os.Getenv("PG_URL")
	if connStr == "" {
		log.Fatal("can't read database connection string from `PG_URL` environment var")
	}

	s, err := store.New(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
}
