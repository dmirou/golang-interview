package shiftingletters

import "testing"

func TestShiftingLetters(t *testing.T) {
	t.Log(shiftingLetters("abc", []int{3, 5, 9}))
	t.Log(shiftingLetters("bad", []int{10, 20, 30}))
}
