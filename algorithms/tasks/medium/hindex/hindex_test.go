package hindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "Example 1: [3,0,6,1,5]",
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			name:      "Example 2: [1,3,1]",
			citations: []int{1, 3, 1},
			expected:  1,
		},
		{
			name:      "Single paper with 1 citation",
			citations: []int{1},
			expected:  1,
		},
		{
			name:      "Two papers: [1,2]",
			citations: []int{1, 2},
			expected:  1,
		},
		{
			name:      "Three papers: [1,2,3]",
			citations: []int{1, 2, 3},
			expected:  2,
		},
		{
			name:      "Complex case: [1,2,2,3,4,4,5]",
			citations: []int{1, 2, 2, 3, 4, 4, 5},
			expected:  3,
		},
		{
			name:      "All zeros",
			citations: []int{0, 0, 0, 0},
			expected:  0,
		},
		{
			name:      "Single zero",
			citations: []int{0},
			expected:  0,
		},
		{
			name:      "High citations: [100, 200, 300]",
			citations: []int{100, 200, 300},
			expected:  3,
		},
		{
			name:      "Descending order: [5,4,3,2,1]",
			citations: []int{5, 4, 3, 2, 1},
			expected:  3,
		},
		{
			name:      "Already sorted ascending: [1,2,3,4,5]",
			citations: []int{1, 2, 3, 4, 5},
			expected:  3,
		},
		{
			name:      "Duplicate values: [2,2,2,2]",
			citations: []int{2, 2, 2, 2},
			expected:  2,
		},
		{
			name:      "Edge case: [1,1,1,1,1]",
			citations: []int{1, 1, 1, 1, 1},
			expected:  1,
		},
		{
			name:      "Large numbers: [1000, 1000, 1000]",
			citations: []int{1000, 1000, 1000},
			expected:  3,
		},
		{
			name:      "Mixed zeros and non-zeros: [0,1,3,5,6]",
			citations: []int{0, 1, 3, 5, 6},
			expected:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hIndex(tt.citations)
			assert.Equal(t, tt.expected, result, 
				"hIndex(%v) = %d, expected %d", tt.citations, result, tt.expected)
		})
	}
}

func TestHIndexEdgeCases(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		result := hIndex([]int{})
		assert.Equal(t, 0, result)
	})

	t.Run("Single paper with zero citations", func(t *testing.T) {
		result := hIndex([]int{0})
		assert.Equal(t, 0, result)
	})

	t.Run("Single paper with high citations", func(t *testing.T) {
		result := hIndex([]int{10})
		assert.Equal(t, 1, result)
	})

	t.Run("Two papers with same citations", func(t *testing.T) {
		result := hIndex([]int{2, 2})
		assert.Equal(t, 2, result)
	})

	t.Run("All papers have same citation count", func(t *testing.T) {
		result := hIndex([]int{3, 3, 3, 3, 3})
		assert.Equal(t, 3, result)
	})
}

func TestHIndexPropertyBased(t *testing.T) {
	// Test that h-index is always <= number of papers
	t.Run("h-index <= number of papers", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 4, 5},
			{0, 1, 2, 3, 4},
			{10, 20, 30, 40, 50},
			{1, 1, 1, 1, 1},
		}

		for _, citations := range testCases {
			result := hIndex(citations)
			assert.LessOrEqual(t, result, len(citations), 
				"h-index (%d) should be <= number of papers (%d) for citations %v", 
				result, len(citations), citations)
		}
	})

	// Test that h-index is always >= 0
	t.Run("h-index >= 0", func(t *testing.T) {
		testCases := [][]int{
			{0, 0, 0},
			{1, 2, 3},
			{100, 200, 300},
		}

		for _, citations := range testCases {
			result := hIndex(citations)
			assert.GreaterOrEqual(t, result, 0, 
				"h-index (%d) should be >= 0 for citations %v", result, citations)
		}
	})
}

