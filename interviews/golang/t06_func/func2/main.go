// f started
// panic in f
// defer3, i = 20
// defer2, i = 10
// defer1, i = 20
// result = 25
package main

import "fmt"

func main() {
	fmt.Println("result =", f())
}

func f() (i int) {
	fmt.Println("f started")

	defer func() {
		msg := recover()
		fmt.Println(msg)

		defer func() { fmt.Printf("defer1, i = %d\n", i); i += 5 }()
		defer fmt.Printf("defer2, i = %d\n", i)
		defer func() {
			fmt.Printf("defer3, i = %d\n", i)
		}()

		i = i * 2
	}()

	i = 10

	panic("panic in f")
}
