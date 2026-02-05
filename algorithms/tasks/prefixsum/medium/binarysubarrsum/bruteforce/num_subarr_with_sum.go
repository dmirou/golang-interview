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

1752 ms
*/
func numSubarraysWithSum(nums []int, goal int) int {
	sums := make([]int, len(nums)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := subarrSum(sums, i, j)
			if sum > goal {
				break
			}
			if sum == goal {
				count++
			}
		}
	}
	return count
}

func subarrSum(sums []int, start, end int) int {
	return sums[end+1] - sums[start]
}
