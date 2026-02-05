package longestones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		// Example cases from problem
		{
			name:     "example 1",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			name:     "example 2",
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},

		// Edge cases
		{
			name:     "all ones, k=0",
			nums:     []int{1, 1, 1, 1, 1},
			k:        0,
			expected: 5,
		},
		{
			name:     "all zeros, k=0",
			nums:     []int{0, 0, 0, 0, 0},
			k:        0,
			expected: 0,
		},
		{
			name:     "all zeros, k=2",
			nums:     []int{0, 0, 0, 0, 0},
			k:        2,
			expected: 2,
		},
		{
			name:     "all zeros, k equals length",
			nums:     []int{0, 0, 0, 0, 0},
			k:        5,
			expected: 5,
		},
		{
			name:     "alternating 1s and 0s, k=0",
			nums:     []int{1, 0, 1, 0, 1},
			k:        0,
			expected: 1,
		},
		{
			name:     "alternating 1s and 0s, k=1",
			nums:     []int{1, 0, 1, 0, 1},
			k:        1,
			expected: 3,
		},
		{
			name:     "alternating 1s and 0s, k=2",
			nums:     []int{1, 0, 1, 0, 1},
			k:        2,
			expected: 5,
		},

		// Single element cases
		{
			name:     "single 1, k=0",
			nums:     []int{1},
			k:        0,
			expected: 1,
		},
		{
			name:     "single 0, k=0",
			nums:     []int{0},
			k:        0,
			expected: 0,
		},
		{
			name:     "single 0, k=1",
			nums:     []int{0},
			k:        1,
			expected: 1,
		},

		// Mixed cases
		{
			name:     "k larger than zeros count",
			nums:     []int{1, 0, 0, 1, 1, 0, 1},
			k:        10,
			expected: 7,
		},
		{
			name:     "zeros at beginning and end",
			nums:     []int{0, 0, 1, 1, 1, 0, 0},
			k:        2,
			expected: 5,
		},
		{
			name:     "consecutive zeros in middle",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1},
			k:        2,
			expected: 5,
		},

		// Large k cases
		{
			name:     "k=0 with only ones in middle",
			nums:     []int{0, 1, 1, 1, 0},
			k:        0,
			expected: 3,
		},
		{
			name:     "k=1 with only ones in middle",
			nums:     []int{0, 1, 1, 1, 0},
			k:        1,
			expected: 4,
		},
		{
			name:     "k=2 with only ones in middle",
			nums:     []int{0, 1, 1, 1, 0},
			k:        2,
			expected: 5,
		},

		// Complex pattern
		{
			name:     "complex pattern 1",
			nums:     []int{1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			name:     "complex pattern 2",
			nums:     []int{1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1},
			k:        3,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnes(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result,
				"longestOnes(%v, %d) = %d, expected %d",
				tt.nums, tt.k, result, tt.expected)
		})
	}
}

func TestLongestOnes_BoundaryCases(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		assert.Equal(t, 0, longestOnes([]int{}, 0))
		assert.Equal(t, 0, longestOnes([]int{}, 5))
	})

	t.Run("maximum k", func(t *testing.T) {
		nums := []int{0, 1, 0, 1, 0, 1}
		assert.Equal(t, 6, longestOnes(nums, 10))
	})

	t.Run("all zeros with k=0", func(t *testing.T) {
		nums := []int{0, 0, 0, 0}
		assert.Equal(t, 0, longestOnes(nums, 0))
	})

	t.Run("all ones with large k", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1}
		assert.Equal(t, 5, longestOnes(nums, 10))
	})
}

func TestLongestOnes_PropertyBased(t *testing.T) {
	// Test that result never exceeds array length
	t.Run("result never exceeds length", func(t *testing.T) {
		testCases := []struct {
			nums []int
			k    int
		}{
			{[]int{1, 0, 1, 0}, 0},
			{[]int{1, 0, 1, 0}, 1},
			{[]int{1, 0, 1, 0}, 2},
			{[]int{0, 0, 0, 0}, 2},
			{[]int{1, 1, 1, 1}, 2},
		}

		for _, tc := range testCases {
			result := longestOnes(tc.nums, tc.k)
			assert.LessOrEqual(t, result, len(tc.nums),
				"Result %d should not exceed array length %d for nums=%v, k=%d",
				result, len(tc.nums), tc.nums, tc.k)
		}
	})

	// Test monotonic property: more k should give same or better result
	t.Run("monotonic in k", func(t *testing.T) {
		nums := []int{1, 0, 1, 0, 1, 0, 1}
		prev := -1
		for k := 0; k <= len(nums); k++ {
			current := longestOnes(nums, k)
			if prev != -1 {
				assert.GreaterOrEqual(t, current, prev,
					"With more k (%d -> %d), result should not decrease: %d -> %d",
					k-1, k, prev, current)
			}
			prev = current
		}
	})
}
