package store

import (
	"fmt"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	s, err := New(databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, _ = s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		_ = s.Close()
		return
	}
}
