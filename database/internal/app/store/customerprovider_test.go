package store_test

import (
	"testing"

	"github.com/dmirou/learngo/database/internal/app/model"
	"github.com/dmirou/learngo/database/internal/app/store"
)

func TestCustomerProvider_Add(t *testing.T) {
	testCases := []struct {
		name  string
		email string
		valid bool
	}{
		{
			name:  "valid email",
			email: "user@example.com",
			valid: true,
		},
		{
			name:  "empty email",
			email: "",
			valid: false,
		},
	}

	s, down := store.TestStore(t, databaseURL)
	defer down("customer")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := &model.Customer{Email: tc.email}
			err := s.Customer().Add(c)
			if tc.valid && err != nil {
				t.Fatalf("expected: nil, got: %v", err)
			}
			if !tc.valid && err == nil {
				t.Fatalf("expected: error, got: %v", err)
			}
			if tc.valid && c.ID == 0 {
				t.Fatalf("expected: not empty ID, got: %v", c.ID)
			}
			if !tc.valid && c.ID != 0 {
				t.Fatalf("expected: %d, got: %v", 0, c.ID)
			}
		})
	}
}

func TestCustomerProvider_BatchAdd(t *testing.T) {
	testCases := []struct {
		name   string
		emails []string
		valid  bool
	}{
		//{
		//	name:   "valid email",
		//	emails: []string{"user@example.com", "user2@example.org"},
		//	valid:  true,
		//},
		{
			name:   "one empty email",
			emails: []string{"user2@example.com", "user3@example.org", ""},
			valid:  false,
		},
	}

	s, down := store.TestStore(t, databaseURL)
	defer down()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := make([]*model.Customer, len(tc.emails))
			for i, email := range tc.emails {
				cs[i] = &model.Customer{Email: email}
			}

			err := s.Customer().BatchAdd(cs)
			if tc.valid && err != nil {
				t.Fatalf("expected: nil, got: %v", err)
			}
			if !tc.valid && err == nil {
				t.Fatalf("expected: error, got: %v", err)
			}
		})
	}
}

func TestCustomerProvider_Update(t *testing.T) {
	testCases := []struct {
		name      string
		firstName string
		lastName  string
	}{
		{
			name:      "empty all",
			firstName: "",
			lastName:  "",
		},
		{
			name:      "filled all",
			firstName: "Ivan",
			lastName:  "Ivanov",
		},
		{
			name:      "empty last name",
			firstName: "Ivan",
			lastName:  "",
		},
		{
			name:      "empty first name",
			firstName: "",
			lastName:  "Ivanov",
		},
	}

	s, down := store.TestStore(t, databaseURL)
	defer down("customer")

	c1 := store.TestCustomer()
	err := s.Customer().Add(c1)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c1.FirstName = tc.firstName
			c1.LastName = tc.firstName
			err := s.Customer().Update(c1)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCustomerProvider_Delete(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down("customer")

	c := store.TestCustomer()

	if err := s.Customer().Add(c); err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	if err := s.Customer().Delete(c); err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
	if c.DeleteTime == nil {
		t.Fatalf("expected: not nil, got: %v", c.DeleteTime)
	}
}

func TestCustomerProvider_Find(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down("customer")

	c := store.TestCustomer()

	if err := s.Customer().Add(c); err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	c2, err := s.Customer().Find(c.ID)
	if err != nil {
		t.Fatal(err)
	}

	if c2.ID != c.ID {
		t.Fatalf("expected: %v, got: %v", c.ID, c2.ID)
	}

	notExisting := c.ID + 1

	u, err := s.Customer().Find(notExisting)
	if err != nil {
		t.Fatal(err)
	}
	if u != nil {
		t.Fatalf("expected: %v, got: %v", nil, u)
	}
}

func TestCustomerProvider_FindByEmail(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down("customer")

	c := store.TestCustomer()

	if err := s.Customer().Add(c); err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}

	c2, err := s.Customer().FindByEmail(c.Email)
	if err != nil {
		t.Fatal(err)
	}

	if c2.ID != c.ID {
		t.Fatalf("expected: %v, got: %v", c.ID, c2.ID)
	}
	if c2.Email != c.Email {
		t.Fatalf("expected: %v, got: %v", c.ID, c2.ID)
	}

	notExisting := "not-exist@email.com"

	u, err := s.Customer().FindByEmail(notExisting)
	if err != nil {
		t.Fatal(err)
	}
	if u != nil {
		t.Fatalf("expected: %v, got: %v", nil, u)
	}
}

func TestCustomerProvider_List(t *testing.T) {
	s, down := store.TestStore(t, databaseURL)
	defer down()

	l, err := s.Customer().List()
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 0 {
		t.Fatalf("expected len: %v, got: %v", 0, len(l))
	}
}
