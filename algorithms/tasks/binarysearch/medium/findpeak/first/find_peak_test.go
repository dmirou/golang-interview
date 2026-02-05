package findpeak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(t *testing.T, nums []int, result int)
	}{
		{
			name: "Example 1 - peak at index 2",
			nums: []int{1, 2, 3, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 2, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Example 2 - multiple peaks, can return 1 or 5",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			validate: func(t *testing.T, nums []int, result int) {
				assert.True(t, result == 1 || result == 5, "Result should be 1 or 5, got %d", result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Single element - always a peak",
			nums: []int{1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Two elements - first is peak",
			nums: []int{2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Two elements - second is peak",
			nums: []int{1, 2},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 1, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Three elements - peak at middle",
			nums: []int{1, 3, 2},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 1, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Three elements - peak at beginning",
			nums: []int{3, 2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Three elements - peak at end",
			nums: []int{1, 2, 3},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 2, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Ascending array - peak at end",
			nums: []int{1, 2, 3, 4, 5},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 4, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Descending array - peak at beginning",
			nums: []int{5, 4, 3, 2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Peak in middle",
			nums: []int{1, 2, 1, 2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.True(t, result == 1 || result == 3, "Result should be 1 or 3, got %d", result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Multiple peaks - can return any",
			nums: []int{1, 3, 2, 4, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.True(t, result == 1 || result == 3, "Result should be 1 or 3, got %d", result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Large array - peak in middle",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 9, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Negative numbers - peak at end",
			nums: []int{-5, -4, -3, -2, -1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 4, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Negative numbers - peak at beginning",
			nums: []int{-1, -2, -3, -4, -5},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Mixed positive and negative - peak in middle",
			nums: []int{-3, -2, -1, 0, 1, 0, -1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 4, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Peak at index 0",
			nums: []int{10, 9, 8, 7, 6},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Peak at last index",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 9, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Complex pattern with multiple peaks",
			nums: []int{1, 5, 2, 4, 1, 3, 1},
			validate: func(t *testing.T, nums []int, result int) {
				validPeaks := []int{1, 3, 5}
				assert.Contains(t, validPeaks, result, "Result should be one of %v, got %d", validPeaks, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Large values",
			nums: []int{1000, 2000, 1500, 3000, 2500},
			validate: func(t *testing.T, nums []int, result int) {
				assert.True(t, result == 1 || result == 3, "Result should be 1 or 3, got %d", result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Very large array - peak near end",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 19},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 19, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Very large array - peak near beginning",
			nums: []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			validate: func(t *testing.T, nums []int, result int) {
				assert.Equal(t, 0, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
		{
			name: "Alternating pattern",
			nums: []int{1, 3, 1, 3, 1, 3, 1},
			validate: func(t *testing.T, nums []int, result int) {
				validPeaks := []int{1, 3, 5}
				assert.Contains(t, validPeaks, result, "Result should be one of %v, got %d", validPeaks, result)
				assert.True(t, isPeak(nums, result), "Index %d should be a peak", result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement(tt.nums)
			tt.validate(t, tt.nums, result)
		})
	}
}

// isPeak checks if the element at index i is a peak
func isPeak(nums []int, i int) bool {
	if i < 0 || i >= len(nums) {
		return false
	}

	// Check left neighbor (or -∞ if at beginning)
	leftOK := i == 0 || nums[i] > nums[i-1]

	// Check right neighbor (or -∞ if at end)
	rightOK := i == len(nums)-1 || nums[i] > nums[i+1]

	return leftOK && rightOK
}
