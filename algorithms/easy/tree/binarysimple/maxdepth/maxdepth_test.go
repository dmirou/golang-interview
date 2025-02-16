package maxdepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	root := TreeNode{}

	assert.Equal(t, 1, maxDepth(&root))
}
