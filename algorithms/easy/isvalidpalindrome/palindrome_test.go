package isvalidpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	p := "race a car"
	want := false
	if got := isPalindrome(p); got != want {
		t.Errorf("isPalindrome(%q), want: %v, got: %v", p, want, got)
	}

	p = "A man, a plan, a canal: Panama"
	want = true
	if got := isPalindrome(p); got != want {
		t.Errorf("isPalindrome(%q), want: %v, got: %v", p, want, got)
	}
}
