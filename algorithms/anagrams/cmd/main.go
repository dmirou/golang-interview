package main

import (
	"fmt"

	"github.com/dmirou/learngo/algorithms/anagrams"
)

func main() {
	var s1, s2 string

	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	if anagrams.AreAnagrams(s1, s2) {
		fmt.Println("1")
		return
	}
	fmt.Println("0")
}
