package hindex

import (
	"slices"
)

// len 7
// [1 2 2 3 3 4 5] => 3
func hIndex(citations []int) int {
	slices.Sort(citations)

	for i := len(citations) - 1; i >= 0; i-- {
		c := citations[i]
		// 5 > 6 - 6
		// 4 > 6 - 5
		// 3 >= 6 - 4 - 1
		// 3 >= 6 - 3

		// hIndex := len(citations) - 1 - i

		// [3,0,6,1,5] -> [0,1,3,5,6] len 5
		if c >= len(citations)-2-i {
			return c
		}
	}

	return 0
}
