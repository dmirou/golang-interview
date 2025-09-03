// hello
// 108 l
// 101 e
// 108 l
// 108 l
// 111 o
package main

import "fmt"

func main() {
	str := "hello"
	s := []byte(str)
	s[0] = byte('l')

	fmt.Println(str)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c\n", s[i], s[i])
	}
}
