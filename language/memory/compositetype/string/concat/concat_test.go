package concat

import (
	"strconv"
	"strings"
	"testing"
)

func Strings(count int) []string {
	res := make([]string, count)
	for i := 0; i < count; i++ {
		res[i] = strconv.Itoa(i)
	}

	return res
}

const genCount = 50

func BenchmarkSlowConcat(b *testing.B) {
	elems := Strings(genCount)

	for i := 0; i < b.N; i++ {
		res := slowConcat(elems)
		if res == "" {
			b.Error("result is empty")
		}
	}
}

func BenchmarkBufConcat(b *testing.B) {
	elems := Strings(genCount)

	for i := 0; i < b.N; i++ {
		res := bufConcat(elems)
		if res == "" {
			b.Error("result is empty")
		}
	}
}

func BenchmarkBuilderConcat(b *testing.B) {
	elems := Strings(genCount)

	for i := 0; i < b.N; i++ {
		res := builderConcat(elems)
		if res == "" {
			b.Error("result is empty")
		}
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	elems := Strings(genCount)

	for i := 0; i < b.N; i++ {
		res := strings.Join(elems, "")
		if res == "" {
			b.Error("result is empty")
		}
	}
}
