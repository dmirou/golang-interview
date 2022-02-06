package store_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("PG_TEST_DB_URL")
	if databaseURL == "" {
		panic("PG_TEST_DB_URL env var is empty")
	}

	os.Exit(m.Run())
}
