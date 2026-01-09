package decryptmd5

import "testing"

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}
