package main

import "fmt"

func generate(s string, open, closed, n int) {
	if len(s) == 2*n {
		fmt.Println(s)
		return
	}
	if open < n {
		generate(s+"(", open+1, closed, n)
	}
	if closed < open {
		generate(s+")", open, closed+1, n)
	}
}

func main() {
	fmt.Printf("generate(%d)\n", 1)
	generate("", 0, 0, 1)

	fmt.Printf("generate(%d)\n", 2)
	generate("", 0, 0, 2)

	fmt.Printf("generate(%d)\n", 3)
	generate("", 0, 0, 3)
}
