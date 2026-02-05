package largestlesstarget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFloor(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Target exists at first position",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target exists at middle position",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Target exists at last position",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Target exists with duplicates - last occurrence",
			arr:      []int{1, 2, 2, 2, 3, 4, 5},
			target:   2,
			expected: 3,
		},
		{
			name:     "Target exists with multiple duplicates",
			arr:      []int{1, 3, 3, 3, 3, 5, 7},
			target:   3,
			expected: 4,
		},
		{
			name:     "Target smaller than all elements",
			arr:      []int{5, 6, 7, 8, 9},
			target:   3,
			expected: -1,
		},
		{
			name:     "Target larger than all elements - return last element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   10,
			expected: 4,
		},
		{
			name:     "Target between elements - return previous smaller",
			arr:      []int{1, 3, 5, 7, 9},
			target:   4,
			expected: 1,
		},
		{
			name:     "Target between elements - return previous smaller at beginning",
			arr:      []int{1, 3, 5, 7, 9},
			target:   2,
			expected: 0,
		},
		{
			name:     "Target between elements - return previous smaller at end",
			arr:      []int{1, 3, 5, 7, 9},
			target:   8,
			expected: 3,
		},
		{
			name:     "Single element array - target exists",
			arr:      []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element array - target smaller",
			arr:      []int{5},
			target:   3,
			expected: -1,
		},
		{
			name:     "Single element array - target larger",
			arr:      []int{5},
			target:   7,
			expected: 0,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Two elements - target at first",
			arr:      []int{2, 4},
			target:   2,
			expected: 0,
		},
		{
			name:     "Two elements - target at second",
			arr:      []int{2, 4},
			target:   4,
			expected: 1,
		},
		{
			name:     "Two elements - target between - return first",
			arr:      []int{2, 4},
			target:   3,
			expected: 0,
		},
		{
			name:     "Two elements - target smaller",
			arr:      []int{2, 4},
			target:   1,
			expected: -1,
		},
		{
			name:     "Two elements - target larger - return second",
			arr:      []int{2, 4},
			target:   5,
			expected: 1,
		},
		{
			name:     "All elements same - target exists",
			arr:      []int{5, 5, 5, 5, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "All elements same - target smaller",
			arr:      []int{5, 5, 5, 5, 5},
			target:   3,
			expected: -1,
		},
		{
			name:     "All elements same - target larger",
			arr:      []int{5, 5, 5, 5, 5},
			target:   7,
			expected: 4,
		},
		{
			name:     "Negative numbers - target exists",
			arr:      []int{-5, -3, -1, 1, 3, 5},
			target:   -1,
			expected: 2,
		},
		{
			name:     "Negative numbers - target smaller",
			arr:      []int{-5, -3, -1, 1, 3, 5},
			target:   -7,
			expected: -1,
		},
		{
			name:     "Negative target in positive array",
			arr:      []int{1, 3, 5, 7},
			target:   -2,
			expected: -1,
		},
		{
			name:     "Target zero",
			arr:      []int{-3, -1, 0, 2, 4},
			target:   0,
			expected: 2,
		},
		{
			name:     "Target zero - return last negative",
			arr:      []int{-3, -1, 2, 4},
			target:   0,
			expected: 1,
		},
		{
			name:     "Large array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:   8,
			expected: 7,
		},
		{
			name:     "Target just below existing element",
			arr:      []int{1, 3, 5, 7, 9},
			target:   2,
			expected: 0,
		},
		{
			name:     "Target just above existing element",
			arr:      []int{1, 3, 5, 7, 9},
			target:   6,
			expected: 2,
		},
		{
			name:     "Gaps in array - target in gap",
			arr:      []int{10, 20, 30, 40, 50},
			target:   25,
			expected: 1,
		},
		{
			name:     "Duplicates - target equals last duplicate",
			arr:      []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: 3,
		},
		{
			name:     "Target equals minimum element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target smaller than minimum",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "Target equals maximum element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Target larger than maximum",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: 4,
		},
		{
			name:     "Consecutive duplicates at beginning",
			arr:      []int{1, 1, 1, 2, 3, 4},
			target:   1,
			expected: 2,
		},
		{
			name:     "Consecutive duplicates at end",
			arr:      []int{1, 2, 3, 4, 4, 4},
			target:   4,
			expected: 5,
		},
		{
			name:     "Target between two consecutive duplicates",
			arr:      []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: 3,
		},
		{
			name:     "Target just above minimum",
			arr:      []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target just below maximum",
			arr:      []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFloor(tt.arr, tt.target)
			assert.Equal(t, tt.expected, result, "findFloor(%v, %d) = %d, want %d", tt.arr, tt.target, result, tt.expected)
		})
	}
}
