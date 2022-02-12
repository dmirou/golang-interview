package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	c := make(Sequence, 0, len(s))
	return append(c, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main() {
	seq := make(Sequence, 0)
	seq = append(seq, 5, 2, 10, 46, 1)
	fmt.Printf("seq is %s", seq)
}
