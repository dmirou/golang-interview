package subarrprodlessk

// numSubarrayProductLessThanK counts the number of contiguous subarrays
// where the product of all elements is strictly less than k.
// Uses prefix product approach with binary search.
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	n := len(nums)
	// Build prefix product array: prs[i] = product of nums[0] to nums[i-1]
	prs := make([]int, n+1)
	prs[0] = 1
	for i := 1; i <= n; i++ {
		prs[i] = prs[i-1] * nums[i-1]
	}

	count := 0
	// For each starting index i, find the rightmost index j such that
	// product of subarray nums[i:j+1] < k
	// Product of nums[i:j+1] = prs[j+1] / prs[i]
	// We need: prs[j+1] / prs[i] < k  =>  prs[j+1] < k * prs[i]
	for i := 0; i < n; i++ {
		maxVal := k * prs[i]

		// Binary search for the rightmost position where prs[j+1] < maxVal
		left := i + 1
		right := n
		rightmost := i

		for left <= right {
			mid := left + (right-left)/2
			if prs[mid] < maxVal {
				rightmost = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		// Count all subarrays starting at i: from i to rightmost-1
		if rightmost > i {
			count += rightmost - i
		}
	}

	return count
}
