// Infinite loop.
// The same for:
//
//	for i := 0; i < slice[0]; i++ {
//			slice[0]++
//			fmt.Println("Элемент увеличен")
//		}
package main

import "fmt"

func main() {
	slice := []int{1, 2}

	for i := 0; i < len(slice); i++ {
		slice = append(slice, 1)
		fmt.Println("Элемент добавлен в слайс")
	}

	fmt.Println(slice)
}
