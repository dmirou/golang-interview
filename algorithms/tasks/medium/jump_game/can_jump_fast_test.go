package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJumpFast(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1: [2,3,1,1,4] should return true",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2: [3,2,1,0,4] should return false",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Example 3: [1] should return true",
			nums:     []int{1},
			expected: true,
		},
		{
			name:     "Edge case: empty array should return true",
			nums:     []int{},
			expected: true,
		},
		{
			name:     "Edge case: single element [0] should return true",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Edge case: [0,1] should return false",
			nums:     []int{0, 1},
			expected: false,
		},
		{
			name:     "Edge case: [1,0,1,0] should return false",
			nums:     []int{1, 0, 1, 0},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJumpFast(tt.nums)
			assert.Equal(t, tt.expected, result, "canJumpFast(%v) should return %v", tt.nums, tt.expected)
		})
	}
}
