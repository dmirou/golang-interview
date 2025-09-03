package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(1<<10, 1024)                // 1kb 			1024
	fmt.Println(1<<20, 1024*1024)           // 1mb 	  	 1048576
	fmt.Println(1<<30, 1024*1024*1024)      // 1gb 	  1073741824
	fmt.Println(1<<40, 1024*1024*1024*1024) // 1tb 1099511627776

	fmt.Println(10 + 01 + 010)   // 19 = 10 + 1 + 8
	fmt.Println(10 + 0x1 + 0x10) // 27 = 10 + 1 + 16
	fmt.Println(10 + 0b1 + 0b10) // 13 = 10 + 1 + 2
	fmt.Println(0b1000 == 010)   // true, 8 == 8

	var a uint8 = math.MaxUint8     // 255 = 1<<8 - 1
	fmt.Printf("a = %08b\n", a)     // a = 11111111 (255)
	fmt.Printf("a+1 = %08b\n", a+1) // a = 00000000 (0)
	fmt.Printf("a+2 = %08b\n", a+2) // a = 00000001 (1)

	var b int8 = math.MaxInt8                           // 1<<7 - 1  // 127
	fmt.Printf("b = %08b, (b+1)[10]: %d\n", b, b)       // b = 01111111 (127)
	fmt.Printf("b+1 = %09b, (b+1)[10]: %d\n", b+1, b+1) // b+1 = -10000000 (-128)
	fmt.Printf("b+2 = %09b, (b+2)[10]: %d\n", b+2, b+2) // b+2 = -01111111 (-127)
	fmt.Printf("b+3 = %09b, (b+3)[10]: %d\n", b+3, b+3) // b+3 = -01111110 (-126)
}
