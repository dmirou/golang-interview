package main

import "fmt"

// it will change slice with pointer p outside the function
func Append1(p *[]byte, data []byte) {
	*p = append(*p, data...)
}

// it won't change slice with pointer p outside the function
func Append2(p *[]byte, data []byte) {
	sl := append(*p, data...)
	p = &sl
}

func main() {
	bs := make([]byte, 0)

	fmt.Printf("before Append1 %v, pointer is %p\n", bs, bs)
	Append1(&bs, []byte{'a', 'b', 'c'})
	fmt.Printf("after Append1 %v, pointer is %p\n", bs, bs)

	Append2(&bs, []byte{'d', 'e', 'f'})
	fmt.Printf("after Append2 %v, pointer is %p\n", bs, bs)
}
