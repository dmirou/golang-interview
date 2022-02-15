package main

import (
	"fmt"
)

// 1,2,3,4,5 <- исходный
// 5 4 3 2 5 <- результат
func main() {
	a := []int{1, 2, 3, 4, 5}
	a[1] = 4
	modify(a)
	fmt.Println(a)
}

func modify(a []int) {
	a[0] = 5
	a[3] = 2
	a = append(a, 0)
	a[4] = 1
}
