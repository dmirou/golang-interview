package checksubarrsum

/*
A good subarray is a subarray where:

its length is at least two, and
the sum of the elements of the subarray is a multiple of k.

Note that:
A subarray is a contiguous part of the array.
An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

Example 1:

Input: nums = [23,2,4,6,7], k = 6
Output: true
Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

Example 2:
Input: nums = [23,2,6,4,7], k = 6
Output: true
Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

Example 3:
Input: nums = [23,2,6,4,7], k = 13
Output: false

Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 109
0 <= sum(nums[i]) <= 231 - 1
1 <= k <= 231 - 1
*/
func checkSubarraySum(nums []int, k int) bool {
	/*
		check if there is a subarray where:
		- its length is at least two, and
		- the sum of the elements of the subarray is a multiple of k.

		(sums[j+1]-sums[i])%k == 0
		sums[j+1] % k == sums[i] % k

		nums 		[23, 2, 4, 6, 7], k = 6
		reminders 	[0,  5, 1, 5, 5, 0]
		output true [2,4]...

		nums 		[23, 2, 6, 4, 7], k = 6
		reminders 	[0,  5, 1, 1, 5, 0]
		output true [2,6,4]

		nums 		[23, 2,  6, 4, 7], k = 13
		reminders 	[0, 10, 12, 5, 9, 3]
		output false

		nums 		[5, 0, 0, 0] k =3
		reminders	[0, 2, 2, 2, 2]
	*/
	remainders := make([]int, len(nums)+1)
	for i := 1; i < len(remainders); i++ {
		remainders[i] = (remainders[i-1] + nums[i-1]) % k
	}
	remainderIndex := make(map[int]int)
	for i := 0; i < len(remainders); i++ {
		rem := remainders[i]
		if prevIndex, ok := remainderIndex[rem]; ok {
			if i-prevIndex >= 2 {
				return true
			}
			continue
		}
		remainderIndex[rem] = i
	}

	return false
}

/*
nums = [23,2,4,6,6] k = 7
i 0, sum 23%7=2, {0: -1, 2:0}
i 1, sum (2+2)%7=4, {0: -1, 2:0, 4:1}
i 2, sum (4+4)%7=1, {0: -1, 2:0, 4:1, 1:2}
i 3, sum (1+6)%7=0, {0: -1, 2:0, 4:1, 1:2, 0:3}
i 4, sum (0+6)%7=6, {0: -1, 2:0, 4:1, 1:2, 0:3, 6:4}
*/
func checkSubarraySumOptimized(nums []int, k int) bool {
	sum := 0
	remainderIndex := make(map[int]int)
	remainderIndex[0] = -1 // Initialize to handle subarrays starting at index 0

	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		if prevIndex, ok := remainderIndex[sum]; ok {
			if i-prevIndex >= 2 {
				return true
			}
			continue
		}
		remainderIndex[sum] = i
	}

	return false
}
