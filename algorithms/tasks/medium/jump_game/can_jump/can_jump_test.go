package can_jump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			assert.Equal(t, tt.expected, result, "canJump(%v) should return %v", tt.nums, tt.expected)
		})
	}
}
