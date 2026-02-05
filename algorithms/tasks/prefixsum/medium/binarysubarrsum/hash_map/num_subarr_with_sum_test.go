package numsubarrwithsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		goal     int
		expected int
	}{
		// Provided examples from the problem
		{
			name:     "Example 1 - [1,0,1,0,1] with goal 2",
			nums:     []int{1, 0, 1, 0, 1},
			goal:     2,
			expected: 4,
		},
		{
			name:     "Example 2 - [0,0,0,0,0] with goal 0",
			nums:     []int{0, 0, 0, 0, 0},
			goal:     0,
			expected: 15,
		},

		// Edge cases
		{
			name:     "Empty array (shouldn't happen per constraints, but good to test)",
			nums:     []int{},
			goal:     0,
			expected: 0,
		},
		{
			name:     "Single element 0 with goal 0",
			nums:     []int{0},
			goal:     0,
			expected: 1,
		},
		{
			name:     "Single element 1 with goal 1",
			nums:     []int{1},
			goal:     1,
			expected: 1,
		},
		{
			name:     "Single element 1 with goal 0",
			nums:     []int{1},
			goal:     0,
			expected: 0,
		},
		{
			name:     "All ones array with goal equal to length",
			nums:     []int{1, 1, 1},
			goal:     3,
			expected: 1,
		},
		{
			name:     "All ones array with goal 2",
			nums:     []int{1, 1, 1},
			goal:     2,
			expected: 2, // [1,1] (first two), [1,1] (last two)
		},

		// Mixed cases
		{
			name:     "Alternating 1s and 0s",
			nums:     []int{1, 0, 1, 0, 1, 0},
			goal:     2,
			expected: 6,
		},
		{
			name:     "Consecutive 1s at start",
			nums:     []int{1, 1, 0, 0, 1},
			goal:     2,
			expected: 4,
		},
		{
			name:     "Goal 0 with mixed array",
			nums:     []int{1, 0, 0, 1, 0},
			goal:     0,
			expected: 4,
		},
		{
			name:     "Longer run of zeros with goal 0",
			nums:     []int{0, 0, 0},
			goal:     0,
			expected: 6,
		},

		// Boundary cases
		{
			name:     "Goal larger than possible sum",
			nums:     []int{1, 0, 1},
			goal:     5,
			expected: 0,
		},
		{
			name:     "Goal 0 with all ones",
			nums:     []int{1, 1, 1},
			goal:     0,
			expected: 0,
		},

		// Complex cases
		{
			name:     "Multiple subarrays with same sum",
			nums:     []int{0, 1, 0, 1, 0, 1, 0},
			goal:     2,
			expected: 8,
		},
		{
			name:     "Array starting and ending with 1",
			nums:     []int{1, 0, 0, 1, 0, 1},
			goal:     2,
			expected: 5,
		},

		// Large array simulation (not too large for tests)
		{
			name:     "Pattern of zeros and ones",
			nums:     []int{1, 0, 0, 1, 0, 1, 1, 0, 0, 0},
			goal:     2,
			expected: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSubarraysWithSum(tt.nums, tt.goal)
			assert.Equal(t, tt.expected, result,
				"numSubarraysWithSum(%v, %d) = %d, expected %d",
				tt.nums, tt.goal, result, tt.expected)
		})
	}
}

func TestNumSubarraysWithSum_Performance(t *testing.T) {
	// Test with maximum constraints to ensure performance
	// Note: Using smaller size for unit tests
	t.Run("Large array with many zeros", func(t *testing.T) {
		// Create an array of 1000 zeros
		nums := make([]int, 1000)
		goal := 0
		// Expected result for all zeros: n*(n+1)/2
		expected := 1000 * 1001 / 2

		result := numSubarraysWithSum(nums, goal)
		assert.Equal(t, expected, result)
	})

	t.Run("Large array with alternating pattern", func(t *testing.T) {
		// Create alternating pattern
		nums := make([]int, 200)
		for i := range nums {
			nums[i] = i % 2 // Creates [0,1,0,1,0,1,...]
		}
		goal := 25
		// We're testing that it completes without timing out
		result := numSubarraysWithSum(nums, goal)
		assert.True(t, result >= 0, "Result should be non-negative")
	})
}

func TestNumSubarraysWithSum_SpecificPatterns(t *testing.T) {
	// Additional specific pattern tests
	tests := []struct {
		name     string
		nums     []int
		goal     int
		expected int
	}{
		{
			name:     "Multiple overlapping subarrays",
			nums:     []int{1, 0, 1, 0, 1, 0, 1},
			goal:     2,
			expected: 8,
		},
		{
			name:     "All ones with various goals",
			nums:     []int{1, 1, 1, 1},
			goal:     2,
			expected: 3, // [1,1] positions: (0,1), (1,2), (2,3)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSubarraysWithSum(tt.nums, tt.goal)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function to generate test cases for property-based testing
func TestNumSubarraysWithSum_PropertyBased(t *testing.T) {
	// Test the property that sum of counts for all goals should equal total subarrays
	t.Run("Sum of all goals equals total subarrays", func(t *testing.T) {
		nums := []int{1, 0, 1, 0, 1, 0, 1}
		n := len(nums)
		totalSubarrays := n * (n + 1) / 2

		totalFromGoals := 0
		for goal := 0; goal <= n; goal++ {
			totalFromGoals += numSubarraysWithSum(nums, goal)
		}

		assert.Equal(t, totalSubarrays, totalFromGoals,
			"Sum of counts for all goals should equal total number of subarrays")
	})
}
