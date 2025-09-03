// ивет 8 (и в е т) 2 bytes for each cyrillic symbol
//
// 0 1080
// 2 1074
// 4 1077
// 6 1090
//
// Question: How to get rune counts in string?
//
//	utf8.RuneCountInString(substr)
//
// Question: How to print all string chars using for i cycle
//
//	for i := 0; i < len(str); {
//		r, size := utf8.DecodeRune([]byte(str[i]))
//		fmt.Println(string(r), size)
//		i += size
//	}
//	 or
//	rs := []rune(substr)
//	for i := 0; i < len(rs); i++ {
//		fmt.Printf("%d %c\n", rs[i], rs[i])
//	}
package main

import (
	"fmt"
)

func main() {
	var str = "привет"
	substr := str[4:]

	fmt.Println(substr, len(substr)) // 1

	for a, b := range substr {
		fmt.Println(a, b) // 2
	}
}
