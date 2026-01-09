package strrotate_test

import (
	"testing"
)

var testCases = []struct {
	name string
	in   string
	out  string
}{
	{
		"empty",
		"",
		"",
	},
	{
		"one char ASCII",
		"a",
		"a",
	},
	{
		"one char UTF",
		"Я",
		"Я",
	},
	{
		"ASCII string",
		"123abcdefg56",
		"65gfedcba321",
	},
	{
		"UTF string",
		"12привет мир 345",
		"543 рим тевирп21",
	},
}

func TestRotate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Rotate(tc.in)
			if tc.out != actual {
				t.Fatalf("expected: %s, got: %s", tc.out, actual)
			}
		})
	}
}

func TestRotate2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Rotate2(tc.in)
			if tc.out != actual {
				t.Fatalf("expected: %s, got: %s", tc.out, actual)
			}
		})
	}
}
