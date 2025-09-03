// hello
// aello world
// aello world
package main

import (
	"fmt"
	"unsafe"
)

func changeString(s string) []byte {
	startPtr := unsafe.StringData(s)
	sl := unsafe.Slice(startPtr, len(s))
	sl[0] = byte('a')
	return sl
}

func main() {
	s := "hello"
	s2 := s + " world"

	sl := changeString(s2)

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(string(sl))
}
