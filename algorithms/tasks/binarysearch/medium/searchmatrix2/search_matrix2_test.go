package searchmatrix2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	assert.Equal(t, true, searchMatrix([][]int{
		{-1, 3},
	}, -1))

	assert.Equal(t, true, searchMatrix([][]int{
		{1, 4, 7, 11, 1},
		{2, 5, 8, 12, 1},
		{3, 6, 9, 16, 2},
		{10, 13, 14, 17, 2},
		{18, 21, 23, 26, 3},
	}, 5))
}
