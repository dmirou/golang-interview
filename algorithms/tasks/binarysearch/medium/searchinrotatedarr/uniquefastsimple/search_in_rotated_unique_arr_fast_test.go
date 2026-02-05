package uniquefastsimple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1 - target in right half",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Example 2 - target not found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Example 3 - single element not found",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "Single element found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "No rotation - target at beginning",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			target:   0,
			expected: 0,
		},
		{
			name:     "No rotation - target at middle",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			target:   4,
			expected: 3,
		},
		{
			name:     "No rotation - target at end",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			target:   7,
			expected: 6,
		},
		{
			name:     "No rotation - target not found",
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			target:   3,
			expected: -1,
		},
		{
			name:     "Rotated - target in left half",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "Rotated - target in left half at beginning",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
		{
			name:     "Rotated - target in left half at end",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			name:     "Rotated - target in right half at beginning",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Rotated - target in right half at end",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 6,
		},
		{
			name:     "Rotated by 1",
			nums:     []int{7, 0, 1, 2, 4, 5, 6},
			target:   0,
			expected: 1,
		},
		{
			name:     "Rotated by 1 - target in left",
			nums:     []int{7, 0, 1, 2, 4, 5, 6},
			target:   7,
			expected: 0,
		},
		{
			name:     "Rotated by 1 - target in right",
			nums:     []int{7, 0, 1, 2, 4, 5, 6},
			target:   4,
			expected: 4,
		},
		{
			name:     "Rotated by 2",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   0,
			expected: 2,
		},
		{
			name:     "Rotated by 2 - target in left",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   7,
			expected: 1,
		},
		{
			name:     "Rotated by 2 - target in right",
			nums:     []int{6, 7, 0, 1, 2, 4, 5},
			target:   4,
			expected: 5,
		},
		{
			name:     "Two elements - target at first",
			nums:     []int{1, 0},
			target:   1,
			expected: 0,
		},
		{
			name:     "Two elements - target at second",
			nums:     []int{1, 0},
			target:   0,
			expected: 1,
		},
		{
			name:     "Two elements - target not found",
			nums:     []int{1, 0},
			target:   2,
			expected: -1,
		},
		{
			name:     "Two elements - no rotation",
			nums:     []int{0, 1},
			target:   0,
			expected: 0,
		},
		{
			name:     "Two elements - no rotation - target at second",
			nums:     []int{0, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "Three elements - rotated",
			nums:     []int{2, 0, 1},
			target:   0,
			expected: 1,
		},
		{
			name:     "Three elements - rotated - target at first",
			nums:     []int{2, 0, 1},
			target:   2,
			expected: 0,
		},
		{
			name:     "Three elements - rotated - target at last",
			nums:     []int{2, 0, 1},
			target:   1,
			expected: 2,
		},
		{
			name:     "Three elements - no rotation",
			nums:     []int{0, 1, 2},
			target:   1,
			expected: 1,
		},
		{
			name:     "Target at pivot point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			name:     "Target just after pivot",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Target just before pivot",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			name:     "Large rotation - target in left",
			nums:     []int{5, 6, 7, 0, 1, 2, 3, 4},
			target:   6,
			expected: 1,
		},
		{
			name:     "Large rotation - target in right",
			nums:     []int{5, 6, 7, 0, 1, 2, 3, 4},
			target:   2,
			expected: 5,
		},
		{
			name:     "Target smaller than all",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   -1,
			expected: -1,
		},
		{
			name:     "Target larger than all",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   10,
			expected: -1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-5, -3, -1, 0, 1, 3, 5},
			target:   -1,
			expected: 2,
		},
		{
			name:     "Negative numbers - rotated",
			nums:     []int{0, 1, 3, 5, -5, -3, -1},
			target:   -1,
			expected: 6,
		},
		{
			name:     "Negative numbers - rotated - target in left",
			nums:     []int{0, 1, 3, 5, -5, -3, -1},
			target:   3,
			expected: 2,
		},
		{
			name:     "Target between elements in left half",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "Target between elements in right half",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		{
			name:     "Rotated at middle",
			nums:     []int{3, 4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 5,
		},
		{
			name:     "Rotated at middle - target in left",
			nums:     []int{3, 4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 2,
		},
		{
			name:     "Rotated at middle - target in right",
			nums:     []int{3, 4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result, "search(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
		})
	}
}
