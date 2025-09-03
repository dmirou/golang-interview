// Проп
// ��
// [10] 1
// [92 110] 2
package main

import "fmt"

func main() {
	data := []rune{1055, 1088, 1086, 1087}
	fmt.Println(string(data)) // 1

	// not utf-8 symbols
	data = []rune{0xFFFFFFF, 0xFFFFFFA}
	fmt.Println(string(data)) // 2

	data1 := []byte("\n")
	data2 := []byte(`\n`)
	fmt.Println(data1, len(data1)) // 3
	fmt.Println(data2, len(data2)) // 4
}
