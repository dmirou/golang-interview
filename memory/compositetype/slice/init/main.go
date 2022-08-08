package main

import (
	"bytes"
	"fmt"
)

func printBytes(name string, b []byte) {
	fmt.Printf("%s %T, p: %p, val: %v, str: %q, len: %d, cap: %d\n",
		name, b, b, b, b, len(b), cap(b))
}

func main() {

	var s01 []byte // nil
	printBytes("s01", s01)

	s02 := []byte(nil) //nil
	printBytes("s02", s02)

	// reference to an empty struct, global variable in the runtime
	// we can use empty slice instead of nil slice when our func returns nil error
	// and results are empty
	var s03 = []byte{}
	printBytes("s03", s03)

	var emptyStruct struct{}
	fmt.Printf("emptyStruct %T, p: %p, val: %v\n", emptyStruct, &emptyStruct, emptyStruct)

	var base = []byte("hello мир")
	printBytes("base", base)

	sIdx := bytes.Index(base, []byte(" "))

	s1 := base[sIdx+1:]
	printBytes("s1", s1)

	s2 := base[:sIdx]
	printBytes("s2", s2)

	s3 := s2[:len(s2)+3]
	printBytes("s3", s3)

	s3 = append(s3, []byte("иша")...)
	printBytes("s3", s3)

	var s4 = make([]byte, len(s3))
	copy(s4, s3)
	printBytes("s4", s4)

	var s5 []byte
	for _, v := range s4 {
		printBytes("s5", s5)
		s5 = append(s5, v)
	}
	printBytes("s5", s5)

}
