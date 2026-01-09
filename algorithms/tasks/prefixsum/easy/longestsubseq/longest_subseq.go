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
func answerQueries(nums []int, queries []int) []int {
	slices.Sort(nums)

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	results := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		for j := len(sums) - 1; j >= 0; j-- {
			if sums[j] <= queries[i] {
				results[i] = j + 1
				break
			}
		}
	}

	return results
}
