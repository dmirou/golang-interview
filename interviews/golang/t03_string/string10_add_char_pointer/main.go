package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "hello"
	ptr := unsafe.StringData(s[4:])
	fmt.Println(ptr, *ptr, string(*ptr)) // 0x509853 111 o
}
