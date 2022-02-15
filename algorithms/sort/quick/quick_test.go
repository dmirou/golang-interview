package quick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		left     int
		right    int
		pivotIdx int
		out      []int
		newIdx   int
	}{
		{
			name:     "one item",
			in:       []int{1},
			left:     0,
			right:    0,
			pivotIdx: 0,
			out:      []int{1},
			newIdx:   0,
		},
		{
			name:     "sorted",
			in:       []int{1, 2, 3, 4, 5},
			left:     0,
			right:    4,
			pivotIdx: 2,
			out:      []int{1, 2, 3, 4, 5},
			newIdx:   2,
		},
		{
			name:     "not sorted all",
			in:       []int{5, 4, 3, 2, 1},
			left:     0,
			right:    4,
			pivotIdx: 2,
			out:      []int{1, 2, 3, 4, 5},
			newIdx:   2,
		},
		{
			name:     "not sorted part",
			in:       []int{5, 4, 3, 2, 1},
			left:     0,
			right:    3,
			pivotIdx: 1,
			out:      []int{2, 3, 4, 5, 1},
			newIdx:   2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newIdx := partition(tc.in, tc.left, tc.right, tc.pivotIdx)
			assert.Equal(t, tc.out, tc.in)
			assert.Equal(t, tc.newIdx, newIdx)
		})
	}
}
