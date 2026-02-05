package findmininrotatedarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1 - rotated 3 times",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Example 2 - rotated 4 times",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Example 3 - rotated 4 times (no rotation visible)",
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Two elements - rotated",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "Two elements - no rotation",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "Three elements - rotated",
			nums:     []int{2, 3, 1},
			expected: 1,
		},
		{
			name:     "Three elements - no rotation",
			nums:     []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "Rotated by 1",
			nums:     []int{5, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "Rotated by 2",
			nums:     []int{4, 5, 1, 2, 3},
			expected: 1,
		},
		{
			name:     "Rotated by 3",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Rotated by 4",
			nums:     []int{2, 3, 4, 5, 1},
			expected: 1,
		},
		{
			name:     "No rotation - sorted array",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			expected: 0,
		},
		{
			name:     "Minimum at beginning after rotation",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Minimum at end before rotation",
			nums:     []int{5, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "Large array - rotated at middle",
			nums:     []int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3},
			expected: 0,
		},
		{
			name:     "Large array - no rotation",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 0,
		},
		{
			name:     "Negative numbers - rotated",
			nums:     []int{-1, 0, 1, 2, -5, -4, -3, -2},
			expected: -5,
		},
		{
			name:     "Negative numbers - no rotation",
			nums:     []int{-5, -4, -3, -2, -1, 0, 1, 2},
			expected: -5,
		},
		{
			name:     "Mixed positive and negative - rotated",
			nums:     []int{3, 4, 5, -2, -1, 0, 1, 2},
			expected: -2,
		},
		{
			name:     "Minimum at pivot point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Rotated exactly n times (appears sorted)",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Minimum is first element",
			nums:     []int{0, 1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "Minimum is last element before rotation",
			nums:     []int{1, 2, 3, 4, 0},
			expected: 0,
		},
		{
			name:     "Large rotation - minimum in right half",
			nums:     []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "Large rotation - minimum in left half",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 0,
		},
		{
			name:     "All same value (edge case - but problem says unique)",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Minimum in middle of array",
			nums:     []int{7, 8, 9, 1, 2, 3, 4, 5, 6},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			assert.Equal(t, tt.expected, result, "findMin(%v) = %d, want %d", tt.nums, result, tt.expected)
		})
	}
}
