package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	isMarried bool  // 1 byte aligned to 8 bytes
	age       int64 // 8 bytes aligned to 8 bytes
	isMale    bool  // 1 byte aligned to 8 bytes
}

type PersonOptimized struct {
	age int64 // 8 bytes aligned to 8 bytes

	isMarried bool // 1 byte, not aligned
	isMale    bool // 1 byte, aligned to 8 bytes together with isMarried
}

func main() {
	var p1 Person

	fmt.Printf("p1 %+v size is %v bytes\n", p1, unsafe.Sizeof(p1))
	fmt.Printf("p1 %+v alignof is %v bytes\n", p1, unsafe.Alignof(p1))
	fmt.Printf("p1 %+v isMarried offest is %v bytes\n", p1, unsafe.Offsetof(p1.isMarried))
	fmt.Printf("p1 %+v age offest is %v bytes\n", p1, unsafe.Offsetof(p1.age))
	fmt.Printf("p1 %+v isMale offest is %v bytes\n", p1, unsafe.Offsetof(p1.isMale))

	var p2 PersonOptimized

	fmt.Printf("p2 %+v size is %v bytes\n", p2, unsafe.Sizeof(p2))
	fmt.Printf("p2 %+v alignof is %v bytes\n", p2, unsafe.Alignof(p2))
	fmt.Printf("p2 %+v age offest is %v bytes\n", p2, unsafe.Offsetof(p2.age))
	fmt.Printf("p2 %+v isMarried offest is %v bytes\n", p2, unsafe.Offsetof(p2.isMarried))
	fmt.Printf("p2 %+v isMale offest is %v bytes\n", p2, unsafe.Offsetof(p2.isMale))
}
