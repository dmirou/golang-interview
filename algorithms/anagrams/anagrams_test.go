package anagrams

import (
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		s1       string
		s2       string
		anagrams bool
	}{
		{
			"empty",
			"",
			"",
			true,
		},
		{
			"same char",
			"a",
			"a",
			true,
		},
		{
			"not same char",
			"a",
			"b",
			false,
		},
		{
			"two chars",
			"ab",
			"ba",
			true,
		},
		{
			"anagram with repeated chars",
			"abacdc",
			"cdbaca",
			true,
		},
		{
			"not the same count of chars in s2",
			"abacdc",
			"cdbac",
			false,
		},
		{
			"not the same count of chars in s1",
			"abcdc",
			"cdbaca",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := AreAnagrams(tc.s1, tc.s2)
			if tc.anagrams != actual {
				t.Errorf("expected: %t, got: %t", tc.anagrams, actual)
			}
		})
	}
}
