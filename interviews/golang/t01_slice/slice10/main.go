// [3, 4] len 2 cap 8
// slice [(3, 4) 5 6 7 8 9 0]
// [4] len 1 cap 2
// slice [(4), _]
package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	t := a[2:4]
	fmt.Println(t, len(t), cap(t)) // ?

	t = a[3:4:5]
	fmt.Println(t, len(t), cap(t)) // ?
}
