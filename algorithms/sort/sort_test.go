package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dmirou/learngo/algorithms/sort/bubble"
)

var testCases = []struct {
	name string
	in   []int
	out  []int
}{
	{
		"nil slice",
		nil,
		nil,
	},
	{
		"empty slice",
		[]int{},
		[]int{},
	},
	{
		"one item",
		[]int{1},
		[]int{1},
	},
	{
		"sorted unique items",
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"sorted not unique items",
		[]int{1, 2, 2, 3, 3, 3, 4, 5, 5, 5},
		[]int{1, 2, 2, 3, 3, 3, 4, 5, 5, 5},
	},
	{
		"reversed unique items",
		[]int{5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"reversed not unique items",
		[]int{5, 4, 3, 3, 2, 1, 1, 1},
		[]int{1, 1, 1, 2, 3, 3, 4, 5},
	},
	{
		"random items",
		[]int{-1, 3, 1, 5, 2, 4, -3, 8, 8, 0, 0, 0, -1, -1, -3, -2},
		[]int{-3, -3, -2, -1, -1, -1, 0, 0, 0, 1, 2, 3, 4, 5, 8, 8},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bubble.Sort(tc.in)
			assert.Equal(t, tc.out, tc.in)
		})
	}
}
