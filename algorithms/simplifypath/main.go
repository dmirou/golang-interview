package main

import (
	"fmt"
	"strings"
)

// Есть путь типа “/home/user1/.././user2/some_folder/./file.txt”
// и надо реализовать метод, который вернёт “/home/user2/some_folder/file.txt”

func simplify(s string) string {
	items := strings.Split(s, "/")

	result := make([]string, 0, len(items))

	for _, itm := range items {
		if itm == "." {
			continue
		}

		if itm != ".." {
			result = append(result, itm)
			continue
		}

		result = result[:len(result)-1]
	}

	return strings.Join(result, "/")
}

func main() {
	in := "/home/user1/.././user2/some_folder/./file.txt"
	fmt.Printf("input: %s\noutput: %s\n", in, simplify(in))
}
