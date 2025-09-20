package min_jumps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpBFS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: [2,3,1,1,4]",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 2: [2,3,0,1,4]",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 3: [2]",
			nums:     []int{2},
			expected: 0,
		},
		{
			name:     "Example 4: [3,1000,1,1,2,0]",
			nums:     []int{3, 1000, 1, 1, 2, 0},
			expected: 2,
		},
		{
			name:     "Example 5: [3,2,3,1,1,0]",
			nums:     []int{3, 2, 3, 1, 1, 0},
			expected: 2,
		},
		{
			name:     "Two elements - direct jump",
			nums:     []int{1, 0},
			expected: 1,
		},
		{
			name:     "Large jump at start",
			nums:     []int{5, 1, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Sequential jumps",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jump_bfs(tt.nums)
			assert.Equal(t, tt.expected, result, "jump(%v) = %d, expected %d", tt.nums, result, tt.expected)
		})
	}
}

func TestJumpBFSEdgeCases(t *testing.T) {
	t.Run("Empty array should panic or return 0", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				// Expected behavior for empty array
				t.Log("Function panicked as expected for empty array")
			}
		}()
		result := jump_greedy([]int{})
		assert.Equal(t, 0, result)
	})

	t.Run("Array with all zeros", func(t *testing.T) {
		// This should not be possible according to constraints, but testing for robustness
		nums := []int{0, 0, 0, 0}
		result := jump_bfs(nums)
		// The function should handle this gracefully
		t.Logf("Result for all zeros: %d", result)
	})
}
