package list_test

import (
	"testing"
)

func TestList_Reverse(t *testing.T) {
	testCases := []struct {
		name     string
		items    []int
		reversed []int
	}{
		{
			name:     "nil",
			items:    nil,
			reversed: nil,
		},
		{
			name:     "one item",
			items:    []int{1},
			reversed: []int{1},
		},
		{
			name:     "two items",
			items:    []int{1, 2},
			reversed: []int{2, 1},
		},
		{
			name:     "three items",
			items:    []int{1, 2, 3},
			reversed: []int{3, 2, 1},
		},
		{
			name:     "five items",
			items:    []int{1, 2, 3, 4, 5},
			reversed: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New()
			for i := len(tc.items) - 1; i >= 0; i-- {
				l.AddToStart(tc.items[i])
			}

			r := l.Reverse()
			rItems := r.GetAll()

			if len(tc.reversed) != len(rItems) {
				t.Fatalf("expected: %d, got: %d\n", len(tc.reversed), len(rItems))
			}

			for i := 0; i < len(tc.reversed); i++ {
				if tc.reversed[i] != rItems[i] {
					t.Fatalf("expected: %d, got: %d\n", tc.reversed[i], rItems[i])
				}
			}
		})
	}
}
