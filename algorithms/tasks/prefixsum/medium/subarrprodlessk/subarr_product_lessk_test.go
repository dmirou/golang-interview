package subarrprodlessk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1: [10,5,2,6], k=100",
			nums:     []int{10, 5, 2, 6},
			k:        100,
			expected: 8,
		},
		{
			name:     "Example 2: [1,2,3], k=0",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: 0,
		},
		{
			name:     "Single element less than k",
			nums:     []int{5},
			k:        10,
			expected: 1,
		},
		{
			name:     "Single element equal to k",
			nums:     []int{10},
			k:        10,
			expected: 0,
		},
		{
			name:     "Single element greater than k",
			nums:     []int{20},
			k:        10,
			expected: 0,
		},
		{
			name:     "All elements less than k",
			nums:     []int{1, 2, 3},
			k:        100,
			expected: 6,
		},
		{
			name:     "No subarrays less than k",
			nums:     []int{10, 20, 30},
			k:        5,
			expected: 0,
		},
		{
			name:     "k=1 edge case",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: 0,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 6,
		},
		{
			name:     "Mixed values",
			nums:     []int{2, 3, 4},
			k:        10,
			expected: 4,
		},
		{
			name:     "Large k value",
			nums:     []int{1, 1, 1, 1},
			k:        1000000,
			expected: 10,
		},
		{
			name:     "Increasing sequence",
			nums:     []int{1, 2, 3, 4, 5},
			k:        20,
			expected: 9,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{5, 4, 3, 2, 1},
			k:        20,
			expected: 9,
		},
		{
			name:     "Two elements both valid",
			nums:     []int{2, 3},
			k:        10,
			expected: 3,
		},
		{
			name:     "Two elements one invalid",
			nums:     []int{2, 6},
			k:        10,
			expected: 2,
		},
		{
			name:     "Repeated same value",
			nums:     []int{3, 3, 3},
			k:        10,
			expected: 5,
		},
		{
			name:     "Repeated same value all valid",
			nums:     []int{2, 2, 2},
			k:        10,
			expected: 6,
		},
		{
			name:     "Case where product equals k",
			nums:     []int{2, 5},
			k:        10,
			expected: 2,
		},
		{
			name:     "Small array with boundary",
			nums:     []int{1, 1, 1, 1, 1},
			k:        1,
			expected: 0,
		},
		{
			name:     "Exact boundary case",
			nums:     []int{10, 5, 2, 6},
			k:        100,
			expected: 8,
		},
		{
			name:     "All products valid except one",
			nums:     []int{1, 1, 10, 1, 1},
			k:        10,
			expected: 6,
		},
		{
			name:     "Binary search edge case",
			nums:     []int{1, 2, 3, 4},
			k:        10,
			expected: 7,
		},
		{
			name:     "Very small k",
			nums:     []int{1, 1, 1},
			k:        1,
			expected: 0,
		},
		{
			name:     "Product exactly at boundary",
			nums:     []int{2, 3, 4},
			k:        24,
			expected: 5,
		},
		{
			name:     "Edge case: single element at boundary",
			nums:     []int{10},
			k:        10,
			expected: 0,
		},
		{
			name:     "Edge case: k=0 with positive numbers",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: 0,
		},
		{
			name:     "Case where all single elements valid but pairs invalid",
			nums:     []int{9, 9},
			k:        10,
			expected: 2,
		},
		{
			name:     "Binary search boundary: finding rightmost valid",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        20,
			expected: 10,
		},
		{
			name:     "Case with product overflow potential",
			nums:     []int{100, 100, 100},
			k:        10000,
			expected: 3,
		},
		{
			name:     "Complex case: mixed valid/invalid",
			nums:     []int{1, 2, 3, 4, 5},
			k:        10,
			expected: 8,
		},
		{
			name:     "Edge case: k=1 with all ones (would cause panic with old code)",
			nums:     []int{1, 1, 1, 1, 1},
			k:        1,
			expected: 0,
		},
		{
			name:     "Edge case: k=1 with larger numbers",
			nums:     []int{2, 2, 2},
			k:        1,
			expected: 0,
		},
		{
			name:     "Single element 333, k=1",
			nums:     []int{333},
			k:        1,
			expected: 0,
		},
		{
			name:     "Large array all > 1, k=1",
			nums:     []int{640, 534, 20, 370, 242, 862, 846, 778, 528, 204, 502, 397, 741, 795, 674, 965, 452, 845, 244, 800},
			k:        1,
			expected: 0,
		},
		{
			name:     "20 threes, k=1000",
			nums:     []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			k:        1000,
			expected: 105,
		},
		{
			name:     "Mixed array, k=1",
			nums:     []int{5, 2, 4, 5, 2, 5, 4, 3, 5, 2, 3, 4, 2, 4, 4, 5, 2, 2, 2, 5},
			k:        1,
			expected: 0,
		},
		{
			name:     "40 elements, k=1000000",
			nums:     []int{9, 4, 8, 1, 5, 6, 3, 9, 1, 8, 2, 3, 7, 8, 9, 8, 4, 4, 5, 7, 5, 7, 10, 7, 9, 2, 1, 7, 4, 6, 6, 8, 9, 2, 2, 10, 6, 2, 6, 8},
			k:        1000000,
			expected: 277,
		},
		{
			name:     "80 elements, k=1000000",
			nums:     []int{5, 2, 1, 2, 2, 6, 2, 7, 6, 5, 2, 4, 1, 2, 8, 2, 3, 9, 6, 9, 8, 5, 4, 8, 4, 4, 9, 10, 7, 9, 1, 4, 4, 2, 8, 8, 8, 2, 4, 2, 5, 1, 5, 2, 3, 2, 1, 3, 9, 6, 2, 1, 6, 2, 10, 4, 5, 3, 4, 5, 5, 10, 8, 5, 2, 8, 9, 6, 10, 8, 5, 5, 2, 7, 2, 6, 8, 2, 7, 10},
			k:        1000000,
			expected: 1397,
		},
		{
			name:     "120 elements, k=1000000",
			nums:     []int{895, 142, 300, 150, 302, 405, 537, 937, 938, 674, 635, 960, 929, 468, 105, 695, 182, 172, 413, 537, 156, 544, 683, 991, 359, 985, 613, 483, 838, 217, 755, 254, 857, 295, 343, 271, 154, 622, 625, 519, 853, 182, 860, 901, 412, 652, 767, 988, 525, 491, 132, 408, 689, 662, 142, 260, 217, 844, 487, 905, 283, 832, 363, 380, 734, 253, 373, 654, 606, 953, 589, 433, 544, 465, 843, 649, 226, 518, 401, 774, 895, 317, 418, 708, 435, 527, 114, 798, 478, 395, 751, 582, 973, 162, 387, 215, 543, 735, 131, 644, 415, 725, 867, 671, 529, 894, 800, 372, 607, 429, 355, 990, 783, 464, 465, 631, 756, 948, 760, 691},
			k:        1000000,
			expected: 3825,
		},
		{
			name: "1000 ones, k=2",
			nums: func() []int {
				ones := make([]int, 1000)
				for i := range ones {
					ones[i] = 1
				}
				return ones
			}(),
			k:        2,
			expected: 500500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSubarrayProductLessThanK(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result, "numSubarrayProductLessThanK(%v, %d) = %d, want %d", tt.nums, tt.k, result, tt.expected)
		})
	}
}
