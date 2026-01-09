package minsubarraylen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSubArrayLen(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 3, minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
