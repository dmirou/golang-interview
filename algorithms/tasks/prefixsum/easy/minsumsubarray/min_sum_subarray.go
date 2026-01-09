package minsumsubarray

/*

Input: nums = [3, -2, 1, 4], l = 2, r = 3
Output: 1

Explanation:
The subarrays of length between l = 2 and r = 3 where the sum is greater than 0 are:

[3, -2] with a sum of 1
[1, 4] with a sum of 5
[3, -2, 1] with a sum of 2
[-2, 1, 4] with a sum of 3
Out of these, the subarray [3, -2] has a sum of 1, which is the smallest positive sum. Hence, the answer is 1.

Input: nums = [-2, 2, -3, 1], l = 2, r = 3
Output: -1

Explanation:
There is no subarray of length between l and r that has a sum greater than 0. So, the answer is -1.

Input: nums = [1, 2, 3, 4], l = 2, r = 4
Output: 3

Explanation:
The subarray [1, 2] has a length of 2 and the minimum sum greater than 0. So, the answer is 3.
*/

func minimumSumSubarray(nums []int, l int, r int) int {
	// 1 <= nums.length <= 100
	// 1 <= l <= r <= nums.length
	// -1000 <= nums[i] <= 1000

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	minSum := -1

	// nums [-2,2,-3,1] l2, r3, len 4
	// sums [-2,0,-3,-2]
	// i0, j2, r1, ls0, s0
	// i0, j3, r2, ls0, s-3
	// i1, j2, r2, ls0, s-3
	// i1, j3, r3, ls0, s-2
	// i2, j2, r3, ls-3, s-2-(-3)=1

	// nums [25, 14] l1, r1
	// sums [25, 39]
	// i0, j1, r0, ls0, s25, minSum25
	// i1, j1, r1, ls25, s14, minSum14

	for i := 0; i < len(sums); i++ {
		for j := l; j <= r; j++ {
			r := i + j - 1
			if r >= len(sums) {
				continue
			}

			ls := 0
			if i >= 1 {
				ls = sums[i-1]
			}

			s := sums[r] - ls
			if r == i && j == 1 {
				s = nums[i]
			}
			if s <= 0 {
				continue
			}
			if minSum == -1 {
				minSum = s
				continue
			}
			if minSum > s {
				minSum = s
			}
		}
	}

	return minSum
}
