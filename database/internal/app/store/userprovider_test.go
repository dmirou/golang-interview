package store_test

import (
	"testing"

	"github.com/dmirou/learngo/database/internal/app/store"
)

func TestUserProvider_Find(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down()

	notExisting := int64(999)

	u, err := s.User().Find(notExisting)
	if err != nil {
		t.Fatal(err)
	}
	if u != nil {
		t.Fatalf("expected: %v, got: %v", nil, u)
	}
}

func TestUserProvider_List(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down()

	l, err := s.User().List()
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 0 {
		t.Fatalf("expected len: %v, got: %v", 0, len(l))
	}
}
