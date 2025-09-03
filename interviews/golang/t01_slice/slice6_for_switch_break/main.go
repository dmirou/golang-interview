// 2
// 4
// 5
// 6
// 6
// ... (infinite loop, because break affects switch)
// How to fix?
// cycle:
//
//	for {
//		switch {
//		default:
//			i++
//		case i > 5:
//			break cycle:
package main

import "fmt"

func main() {
	i := 0
	for {
		switch {
		default:
			i++
		case i > 5:
			break
		case i < 3:
			i += 2
		}
		fmt.Println(i) // ?
	}

	fmt.Println(i) // ?
}
