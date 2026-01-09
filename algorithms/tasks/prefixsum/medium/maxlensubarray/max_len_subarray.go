package maxlensubarray

/*
Example 1:
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

Example 2:
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Example 3:
Input: nums = [0,1,1,1,1,1,0,0,0]
Output: 6
Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.

Constraints:
1 <= nums.length <= 105
nums[i] is either 0 or 1.

Test case 1:
Input nums = [0,1,0,1]
Output 2
Expected 4
*/

func findMaxLength(nums []int) int {
	diff := func(num int) int {
		if num == 1 {
			return 1
		}
		return -1
	}

	sums := make([]int, len(nums))
	sums[0] = diff(nums[0])
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + diff(nums[i])
	}
	/*
		Input
		idx  = [0  1 2 3 4 5 6 7 8]
		nums = [0, 1,1,1,1,1,0,0,0]
		sums = [-1,0,1,2,3,4,3,2,1]
		Output 6
		Expected 6
	*/
	maxLen := 0
	sumIndex := make(map[int]int)

	for i := 0; i < len(sums); i++ {
		if sums[i] == 0 {
			maxLen = max(maxLen, i+1)
			continue
		}
		if idx, ok := sumIndex[sums[i]]; ok {
			maxLen = max(maxLen, i-idx)
			continue
		}
		sumIndex[sums[i]] = i
	}

	return maxLen
}
