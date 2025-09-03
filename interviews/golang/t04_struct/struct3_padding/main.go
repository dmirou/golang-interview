// initial output:
// 32
// {false  [0 0]}
// {false  [0 0]}
// {true abc [1 2]}
//
//	type st struct {
//		p1 bool
//		p2 string
//		p3 [2]byte
//	}
//
// bool uses byte
// string uses 16 bytes (int for pointer, int for len)
// int = int64 = 8 bytes
// 32 (because of padding 1 + 7 padding + 8 str pointer + 8 str len + 2 + 6 padding)
//
// how to fix:
//
//	type st struct {
//		p1 bool
//		p3 [2]byte
//		p2 string
//	}
//
// will need 24 bytes
package main

import (
	"fmt"
	"unsafe"
)

type st struct {
	p1 bool
	p2 string
	p3 [2]byte
}

func main() {
	myStr := st{}

	fmt.Println(unsafe.Sizeof(myStr)) // 1
	fmt.Println(myStr)                // 2
	mutatePtr1(&myStr)
	fmt.Println(myStr) // 3
	mutatePtr2(&myStr)
	fmt.Println(myStr) // 4
}

func mutatePtr1(myStr *st) {
	myStr = &st{
		p1: true,
		p2: "abc",
		p3: [2]byte{1, 2},
	}
}

func mutatePtr2(myStr *st) {
	*myStr = st{
		p1: true,
		p2: "abc",
		p3: [2]byte{1, 2},
	}
}
