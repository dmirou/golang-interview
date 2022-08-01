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
		fmt.Printf("i: %d, rune: %c; ", i, v)
	}

	fmt.Println("\nfor i := 0; i < len(s); i++")
	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d, byte: %d; ", i, s[i])
	}

	fmt.Println("\nfor i, v := range []rune(s)")
	for i, v := range []rune(s) {
		fmt.Printf("i: %d, rune: %c; ", i, v)
	}

	s = "привет hey"
	fmt.Printf(
		"\nvar: %s, value: %s, len: %d, runeCount: %d\n",
		"s", s, len(s), utf8.RuneCountInString(s),
	)

	fmt.Println("for i, v := range s")
	for i, v := range s {
		fmt.Printf("i: %d, rune: %c; ", i, v)
	}

	fmt.Println("\nfor i := 0; i < len(s); i++")
	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d, byte: %d; ", i, s[i])
	}

	fmt.Println("\nfor i, v := range []rune(s)")
	for i, v := range []rune(s) {
		fmt.Printf("i: %d, rune: %c; ", i, v)
	}
}
