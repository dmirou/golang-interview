package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAllValues(t *testing.T) {
	expected := []int{1, 5, 9, -1}
	actual := make([]int, len(expected), len(expected))
	idx := 0

	for v := range SliceAllValues(expected) {
		actual[idx] = v
		idx++
	}

	assert.Equal(t, expected, actual)
}

func TestSliceAllKeyValues(t *testing.T) {
	expected := []int{1, 5, 9, -1}
	actual := make([]int, len(expected), len(expected))

	for idx, v := range SliceAllKeyValues(expected) {
		actual[idx] = v
	}

	assert.Equal(t, expected, actual)
}

func TestSliceAppendSeq(t *testing.T) {
	src := []int{1, 9, 25}
	seq := []int{0, 25, 1}
	want := []int{1, 9, 25, 0, 25, 1}
	actual := SliceAppendSeq(src, SliceAllValues(seq))
	assert.Equal(t, want, actual)
}
