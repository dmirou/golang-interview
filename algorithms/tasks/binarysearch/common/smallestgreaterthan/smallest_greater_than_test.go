package smallestgreaterthan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCeil(t *testing.T) {
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
			name:     "Target exists with duplicates - first occurrence",
			arr:      []int{1, 2, 2, 2, 3, 4, 5},
			target:   2,
			expected: 1,
		},
		{
			name:     "Target exists with multiple duplicates",
			arr:      []int{1, 3, 3, 3, 3, 5, 7},
			target:   3,
			expected: 1,
		},
		{
			name:     "Target smaller than all elements - return first element",
			arr:      []int{5, 6, 7, 8, 9},
			target:   3,
			expected: 0,
		},
		{
			name:     "Target larger than all elements",
			arr:      []int{1, 2, 3, 4, 5},
			target:   10,
			expected: -1,
		},
		{
			name:     "Target between elements - return next larger",
			arr:      []int{1, 3, 5, 7, 9},
			target:   4,
			expected: 2,
		},
		{
			name:     "Target between elements - return next larger at end",
			arr:      []int{1, 3, 5, 7, 9},
			target:   8,
			expected: 4,
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
			expected: 0,
		},
		{
			name:     "Single element array - target larger",
			arr:      []int{5},
			target:   7,
			expected: -1,
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
			name:     "Two elements - target between - return second",
			arr:      []int{2, 4},
			target:   3,
			expected: 1,
		},
		{
			name:     "Two elements - target smaller - return first",
			arr:      []int{2, 4},
			target:   1,
			expected: 0,
		},
		{
			name:     "Two elements - target larger",
			arr:      []int{2, 4},
			target:   5,
			expected: -1,
		},
		{
			name:     "All elements same - target exists",
			arr:      []int{5, 5, 5, 5, 5},
			target:   5,
			expected: 0,
		},
		{
			name:     "All elements same - target smaller",
			arr:      []int{5, 5, 5, 5, 5},
			target:   3,
			expected: 0,
		},
		{
			name:     "All elements same - target larger",
			arr:      []int{5, 5, 5, 5, 5},
			target:   7,
			expected: -1,
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
			expected: 0,
		},
		{
			name:     "Negative target in positive array",
			arr:      []int{1, 3, 5, 7},
			target:   -2,
			expected: 0,
		},
		{
			name:     "Target zero",
			arr:      []int{-3, -1, 0, 2, 4},
			target:   0,
			expected: 2,
		},
		{
			name:     "Target zero - return first positive",
			arr:      []int{-3, -1, 2, 4},
			target:   0,
			expected: 2,
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
			expected: 1,
		},
		{
			name:     "Target just above existing element",
			arr:      []int{1, 3, 5, 7, 9},
			target:   6,
			expected: 3,
		},
		{
			name:     "Gaps in array - target in gap",
			arr:      []int{10, 20, 30, 40, 50},
			target:   25,
			expected: 2,
		},
		{
			name:     "Duplicates - target equals first duplicate",
			arr:      []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: 1,
		},
		{
			name:     "Duplicates - target between duplicates",
			arr:      []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: 1,
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
			expected: -1,
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
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findCeil(tt.arr, tt.target)
			assert.Equal(t, tt.expected, result, "findCeil(%v, %d) = %d, want %d", tt.arr, tt.target, result, tt.expected)
		})
	}
}
