package recursion_memo

import (
	"testing"
)

func TestLargestSumOfAverages(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
		epsilon  float64
	}{
		// Example test cases from problem description
		{
			name:     "Example 1",
			nums:     []int{9, 1, 2, 3, 9},
			k:        3,
			expected: 20.00000,
			epsilon:  1e-6,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        4,
			expected: 20.50000,
			epsilon:  1e-6,
		},

		// Edge cases
		{
			name:     "Single element array, k=1",
			nums:     []int{5},
			k:        1,
			expected: 5.00000,
			epsilon:  1e-6,
		},
		{
			name:     "All same numbers",
			nums:     []int{3, 3, 3, 3},
			k:        3,
			expected: 9.00000,
			epsilon:  1e-6,
		},
		{
			name:     "k equals array length (each element separate)",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: 10.00000, // 1 + 2 + 3 + 4
			epsilon:  1e-6,
		},
		{
			name:     "k equals 1 (entire array as one group)",
			nums:     []int{4, 1, 7, 5, 6, 2},
			k:        1,
			expected: 25.0 / 6.0,
			epsilon:  1e-6,
		},

		// Increasing sequence
		{
			name:     "Strictly increasing sequence",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: 7.5, // [1,2,3,4], [5] → (10/4) + 5 = 2.5 + 5 = 7.5
			epsilon:  1e-6,
		},

		// Decreasing sequence
		{
			name:     "Strictly decreasing sequence",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: 11,
			epsilon:  1e-5,
		},

		// Mixed patterns
		{
			name:     "Peak in middle",
			nums:     []int{1, 2, 10, 2, 1},
			k:        2,        // 13/3 + 3/2 = 4.3 + 1.5
			expected: 5.833333, // [1,2,10], [2,1] → (13/3) + (3/2) = 4.333 + 1.5 = 5.833
			epsilon:  1e-6,
		},

		// k greater than array length should be handled (but constraint says k <= nums.length)
		{
			name:     "Maximum k allowed",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: 10.00000,
			epsilon:  1e-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestSumOfAverages(tt.nums, tt.k)
			diff := abs(result - tt.expected)
			if diff > tt.epsilon {
				t.Errorf("largestSumOfAverages(%v, %d) = %f, want %f (diff: %f, epsilon: %f)",
					tt.nums, tt.k, result, tt.expected, diff, tt.epsilon)
			}
		})
	}
}

// Helper function for absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Additional edge case tests
func TestEdgeCases(t *testing.T) {
	// Test when k is 1 (whole array average)
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{100, 200, 300}, 1},
	}

	for _, tc := range testCases {
		result := largestSumOfAverages(tc.nums, tc.k)
		// Calculate expected: sum of all elements divided by length
		sum := 0
		for _, num := range tc.nums {
			sum += num
		}
		expected := float64(sum) / float64(len(tc.nums))

		if abs(result-expected) > 1e-6 {
			t.Errorf("For k=1, largestSumOfAverages(%v, %d) = %f, want %f",
				tc.nums, tc.k, result, expected)
		}
	}
}

// Test the property that increasing k should never decrease the score
func TestMonotonicProperty(t *testing.T) {
	nums := []int{9, 1, 2, 3, 9, 4, 5, 6, 7, 8}

	// Get results for k from 1 to len(nums)
	results := make([]float64, len(nums))
	for k := 1; k <= len(nums); k++ {
		results[k-1] = largestSumOfAverages(nums, k)
	}

	// Verify monotonic non-decreasing property
	for i := 1; i < len(results); i++ {
		if results[i] < results[i-1]-1e-6 { // Allow for tiny floating point errors
			t.Errorf("Score decreased when k increased: k=%d: %f, k=%d: %f",
				i, results[i-1], i+1, results[i])
		}
	}
}
