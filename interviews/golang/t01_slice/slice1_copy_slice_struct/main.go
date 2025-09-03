// 1 [123 2] [123 2]
// 2 [123 2 32] [123 2]
package main

import "fmt"

func main() {
	var (
		one = []int{1, 2}
		two = one
	)
	two[0] = 123
	fmt.Println(one, two) // 1

	one = append(one, 32)
	fmt.Println(one, two) // 2
}
