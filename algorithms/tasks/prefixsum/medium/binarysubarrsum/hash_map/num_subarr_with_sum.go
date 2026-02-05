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

26ms
*/
func numSubarraysWithSum(nums []int, goal int) int {
	sumsCounts := make(map[int]int)
	sumsCounts[0] = 1
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		target := sum - goal
		count += sumsCounts[target]
		sumsCounts[sum]++
	}
	return count
}

/*
prefixsum[current] - prefixsum[previous] == goal
prefixsum[previous] = prefixsum[current] - goal

prefix[i] - prefix[j] == goal, j < i
prefix[i] - goal == prefix[j]
prefix[current] - goal = prefix[previous] (from map)

Why This Guarantees Contiguity
The reason contiguity is guaranteed is mathematical:
If: prefix[i] = sum(nums[0] to nums[i-1])
And: prefix[j] = sum(nums[0] to nums[j-1])
Where: i > j
Then: prefix[i] - prefix[j] = sum(nums[j] to nums[i-1])
This difference can ONLY represent the sum of a contiguous block from j to i-1.

It's impossible for this to represent a non-contiguous subarray because:
prefix[i] includes everything from start to i-1
prefix[j] includes everything from start to j-1
Their difference removes the first j elements, leaving exactly elements j to i-1
*/
