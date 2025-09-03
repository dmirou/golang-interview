// article https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
package main

import "fmt"

func example1() {
	var a int
	var b, c = &a, &a
	fmt.Println(b, c)   // 0x1040a124 0x1040a124
	fmt.Println(&b, &c) // 0x1040c108 0x1040c110
}

func example2() {
	fn := func(m map[int]int) {
		m = make(map[int]int)
	}

	var m map[int]int
	fn(m)
	fmt.Println(m == nil) // true
}

func example3() {
	fn := func(m *map[int]int) {
		*m = make(map[int]int)
	}

	var m map[int]int
	fn(&m)
	fmt.Println(m == nil) // false
}

// Go does not have pass-by-reference (передача по ссылке) semantics
// because Go does not have reference variables, only pointers.
// In languages like C++ you can declare an alias, or an alternate name
// to an existing variable. This is called a reference variable.
// #include <stdio.h>
//
// int main() {
// int a = 10;
// int &b = a;
// int &c = b;
//
// printf("%p %p %p\n", &a, &b, &c); // 0x7ffe114f0b14 0x7ffe114f0b14 0x7ffe114f0b14
// return 0;
// }
func main() {
	example1()
	example2()
	example3()
}
