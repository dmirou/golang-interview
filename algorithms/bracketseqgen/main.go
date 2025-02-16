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
	var n int
	fmt.Scanf("%d", &n)
	generate("", 0, 0, n)
}
