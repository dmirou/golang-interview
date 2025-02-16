package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixSeq(t *testing.T) {
	assert.Equal(t, -1, fixSeq("((("))
	assert.Equal(t, -1, fixSeq("("))
	assert.Equal(t, 1, fixSeq("(("))
	assert.Equal(t, -1, fixSeq(")("))
	assert.Equal(t, 7, fixSeq("((()))(("))
	assert.Equal(t, 5, fixSeq("(()()("))
	assert.Equal(t, 0, fixSeq(")(()()()))"))
}
