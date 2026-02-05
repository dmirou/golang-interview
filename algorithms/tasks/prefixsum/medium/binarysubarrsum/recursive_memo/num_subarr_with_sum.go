package numsubarrwithsum

/*
930. Binary Subarrays With Sum
Medium
Topics
premium lock icon
Companies

Given a binary array nums and an integer goal, return the number of non-empty
subarrays with a sum goal. A subarray is a contiguous part of the array.

Example 1:
Input: nums = [1,0,1,0,1], goal = 2
Output: 4
Explanation: The 4 subarrays are bolded and underlined below:
[(1,0,1),0,1]
[(1,0,1,0),1]
[1,(0,1,0,1)]
[1,0,(1,0,1)]

Example 2:
Input: nums = [0,0,0,0,0], goal = 0
Output: 15
1 [(0),0,0,0,0]
2 [(0,0),0,0,0]
3 [(0,0,0),0,0]
4 [(0,0,0,0),0]
5 [(0,0,0,0,0)]
6 [0,(0),0,0,0)]
7 [0,(0,0),0,0)]
8 [0,(0,0,0),0)]
9 [0,(0,0,0,0))]
10 [0,0,(0),0,0]
11 [0,0,(0,0),0]
12 [0,0,(0,0,0)]
13 [0,0,0,(0),0]
14 [0,0,0,(0,0)]
15 [0,0,0,0,(0)]

Constraints:
1 <= nums.length <= 3 * 104
nums[i] is either 0 or 1.
0 <= goal <= nums.length

time limit exceeded
*/
func numSubarraysWithSum(nums []int, goal int) int {
	memo := make(map[callParams]int)

	count := 0
	for i := 0; i < len(nums); i++ {
		count += subarrCountRec(nums, memo, i, goal)
	}
	return count
}

type callParams struct {
	start int
	goal  int
}

func subarrCountRec(nums []int, memo map[callParams]int, start, goal int) int {
	if start >= len(nums) {
		return 0
	}

	if nums[start] > goal {
		return 0
	}

	if count, ok := memo[callParams{start, goal}]; ok {
		return count
	}

	count := 0
	if nums[start] == goal {
		count++
	}
	count += subarrCountRec(nums, memo, start+1, goal-nums[start])

	memo[callParams{start, goal}] = count

	return count
}
