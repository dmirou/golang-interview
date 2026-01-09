package longestsubseq

import (
	"slices"
)

/*
Input: nums = [2,3,4,5], queries = [1]
Output: [0]

Input: nums = [4,5,2,1], queries = [3,10,21]
nums = [1 2 4 5]
sums = [1 3 7 12]
queries = [3, 10, 21]
results = [2, 3, 4]
Output: [2,3,4]
*/
func answerQueriesBinarySearch(nums []int, queries []int) []int {
	slices.Sort(nums)

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	results := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		l := 0
		r := len(sums) - 1
		// 1 2 3, 4, 5
		// l 0, r 4, m 2
		// arr[2] 3 < 4
		// l 3, r 4
		// m 3
		// arr[3] 4 == 4

		/*
			sums = [1 3 7 12]
			queries = [3, 10, 21]
			l0, r3,
		*/
		for l <= r {
			mid := l + (r-l)/2
			if sums[mid] == queries[i] {
				results[i] = mid + 1
				break
			}
			if sums[mid] < queries[i] {
				results[i] = mid + 1
				l = mid + 1
				continue
			}
			r = mid - 1
		}
	}

	return results
}
