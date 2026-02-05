package intersection

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			name:     "Example 2",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
		{
			name:     "No intersection",
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Single element intersection",
			nums1:    []int{1},
			nums2:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Single element no intersection",
			nums1:    []int{1},
			nums2:    []int{2},
			expected: []int{},
		},
		{
			name:     "All elements intersect",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "One array is subset",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{2, 4},
			expected: []int{2, 4},
		},
		{
			name:     "Duplicate elements in both arrays",
			nums1:    []int{1, 1, 2, 2, 3},
			nums2:    []int{2, 2, 3, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Empty first array",
			nums1:    []int{},
			nums2:    []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second array",
			nums1:    []int{1, 2, 3},
			nums2:    []int{},
			expected: []int{},
		},
		{
			name:     "Zero values",
			nums1:    []int{0, 1, 2},
			nums2:    []int{0, 3, 4},
			expected: []int{0},
		},
		{
			name:     "Large values within constraints",
			nums1:    []int{1000, 500, 750},
			nums2:    []int{500, 750, 999},
			expected: []int{500, 750},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create copies to avoid modifying original slices
			nums1Copy := make([]int, len(tt.nums1))
			nums2Copy := make([]int, len(tt.nums2))
			copy(nums1Copy, tt.nums1)
			copy(nums2Copy, tt.nums2)

			result := intersection(nums1Copy, nums2Copy)

			// Sort both result and expected for comparison (order doesn't matter)
			slices.Sort(result)
			expectedSorted := make([]int, len(tt.expected))
			copy(expectedSorted, tt.expected)
			slices.Sort(expectedSorted)

			if !slices.Equal(result, expectedSorted) {
				t.Errorf("intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, result, expectedSorted)
			}
		})
	}
}
