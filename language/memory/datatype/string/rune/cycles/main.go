package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "hello"
	fmt.Printf(
		"var: %s, value: %s, len: %d, runeCount: %d\n",
		"s", s, len(s), utf8.RuneCountInString(s),
	)

	fmt.Println("for i, v := range s")
	for i, v := range s {
		fmt.Printf("i: %d, v: %v %c\n", i, v, v)
	}

	fmt.Println("\nfor i := 0; i < len(s); i++")
	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d, s[i]: %v %c\n", i, s[i], s[i])
	}

	fmt.Println("\nfor i, v := range []rune(s)")
	for i, v := range []rune(s) {
		fmt.Printf("i: %d, v: %v %c\n", i, v, v)
	}

	s = "привет hey⬱邛"
	fmt.Printf(
		"\nvar: %s, value: %s, len: %d, runeCount: %d\n",
		"s", s, len(s), utf8.RuneCountInString(s),
	)

	fmt.Println("for i, v := range s")
	for i, v := range s {
		fmt.Printf("i: %d, v: %v %c\n", i, v, v)
	}

	fmt.Println("\nfor i := 0; i < len(s); i++")
	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d, s[i]: %v %c\n", i, s[i], s[i])
	}

	fmt.Println("\nfor i := 0; i < len(s); { decodeRune")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune([]byte(s[i:]))
		fmt.Printf("i: %d, r: %v %c size: %v\n", i, r, r, size)
		i += size
	}

	fmt.Println("\nfor i, v := range []rune(s)")
	for i, v := range []rune(s) {
		fmt.Printf("i: %d, v: %v %c\n", i, v, v)
	}

	fmt.Printf("rune 11057: %c\n", rune(11057))
	fmt.Printf("rune 37019: %c\n", rune(37019))
	fmt.Printf("rune 99999: %c\n", rune(99999))
}
