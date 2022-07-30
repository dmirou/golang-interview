package main

import (
	"fmt"
	"unsafe"
)

type Person struct { // size 24 bytes, alignof 8 bytes
	isMarried bool  // offset 0 bytes, size 1 byte
	age       int64 // offset 8 bytes, size 8 bytes
	isMale    bool  // offset 16 bytes, size 1 byte
}

type PersonOptimized struct { // size 16 bytes, alignof 8 bytes
	age int64 // offset 0 bytes, size 8 bytes

	isMarried bool // offset 8 bytes, size 1 byte
	isMale    bool // offest 9 bytes, size 1 byte
}

type Store struct { // size 48 bytes, alignof is 8 bytes
	name       string // offset 0 bytes, size 16 bytes
	address    string // offset 16 bytes, size 16 bytes
	goodsCount int    // offset 32 bytes, size 8 bytes
	soldCount  int    // offset 40 bytes, size 8 bytes
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

	var s1 Store

	fmt.Printf("s1 %+v size is %v bytes\n", s1, unsafe.Sizeof(s1))
	fmt.Printf("s1.name %+v size is %v bytes\n", s1, unsafe.Sizeof(s1.name))
	fmt.Printf("s1.address %+v size is %v bytes\n", s1, unsafe.Sizeof(s1.address))
	fmt.Printf("s1.goodsCount %+v size is %v bytes\n", s1, unsafe.Sizeof(s1.goodsCount))
	fmt.Printf("s1.soldCount %+v size is %v bytes\n", s1, unsafe.Sizeof(s1.soldCount))
	fmt.Printf("s1 %+v alignof is %v bytes\n", s1, unsafe.Alignof(s1))
	fmt.Printf("s1 %+v name offest is %v bytes\n", s1, unsafe.Offsetof(s1.name))
	fmt.Printf("s1 %+v address offest is %v bytes\n", s1, unsafe.Offsetof(s1.address))
	fmt.Printf("s1 %+v goodsCount offest is %v bytes\n", s1, unsafe.Offsetof(s1.goodsCount))
	fmt.Printf("s1 %+v soldCount offest is %v bytes\n", s1, unsafe.Offsetof(s1.soldCount))
}
