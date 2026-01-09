package jewelty

import "testing"

func BenchmarkCount(b *testing.B) {
	j := "abcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklm"
	s := "abcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgaui"

	for i := 0; i < b.N; i++ {
		Count(j, s)
	}
}

func BenchmarkCountFast(b *testing.B) {
	j := "abcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklmabcdefghijklm"
	s := "abcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgauiabcdefghijklmnabcdefgfsdakjfdjsakfjdskajfkldsjfklasdjfklsdjafskaljfaeiewpeybnnehuhgaui"

	for i := 0; i < b.N; i++ {
		CountFast(j, s)
	}
}
