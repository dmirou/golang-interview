package subarrprodlessk

/*
713. Subarray Product Less Than K

https://leetcode.com/problems/subarray-product-less-than-k/description/?envType=problem-list-v2&envId=prefix-sum

Medium
Topics
premium lock icon
Companies
Hint

Given an array of integers nums and an integer k, return the number of contiguous subarrays where the product
of all the elements in the subarray is strictly less than k.

Example 1:
Input: nums = [10,5,2,6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

Example 2:
Input: nums = [1,2,3], k = 0
Output: 0

Constraints:

1 <= nums.length <= 3 * 104
1 <= nums[i] <= 1000
0 <= k <= 106
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	n := len(nums)
	// Use int64 to avoid overflow in prefix products
	prs := make([]int64, n+1)
	prs[0] = 1
	for i := 1; i <= n; i++ {
		prs[i] = prs[i-1] * int64(nums[i-1])
	}
	// nums [10 5  2  6], k = 100
	// prs [1 10 50 100 600]
	// pr(0,1) = prs[2] / prs[0] = 50 / 1 = 50
	// pr(0,0) = prs[1] / prs[0] = 10 / 1 = 10
	// pr(1, 3) = prs[4] / prs[1] = 600 / 10 = 60
	// pr(i, j) = prs[j+1] / prs[i]
	//
	// pr(i, j) < k
	// prs[j+1] / prs[i] < k
	// prs[j+1] < k * prs[i]

	count := 0
	k64 := int64(k)
	for i := 0; i < n; i++ {
		maxVal := k64 * prs[i]

		// we need to use left as i+1 because, we iterate through products and
		// we need to include ith item into result
		// prs[i] = product of elements from index 0 to i-1 (exclusive)
		// prs[i+1] = product of elements from index 0 to i (inclusive)
		left := i + 1
		right := n
		rightmost := i // rightmost valid index in prs array

		for left <= right {
			mid := left + (right-left)/2
			if prs[mid] < maxVal {
				rightmost = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		// rightmost is the rightmost index in prs where prs[rightmost] < maxVal
		// This means subarray from i to rightmost-1 is valid
		// So j ranges from i to rightmost-1
		if rightmost > i {
			j := rightmost - 1
			count += j - i + 1
		}
	}

	return count
}
