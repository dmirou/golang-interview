// 1 true
// 2 false
// 3 false
// 4 false
// 5 false
// 6 true
//
// Under the hood, interface values can be thought of as a
// tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name
// on its underlying type.
package main

import "fmt"

type errorCustom struct {
	message string
}

func (e errorCustom) Error() string {
	return e.message
}

func isNil(err error) bool {
	return err == nil
}

func main() {
	var err1 error
	fmt.Println(isNil(err1)) // 1

	var err2 *errorCustom
	fmt.Println(isNil(err2)) // 2

	err2 = &errorCustom{"error"}
	fmt.Println(isNil(err2)) // 3

	err2 = nil
	fmt.Println(isNil(err2)) // 4

	err1 = err2
	fmt.Println(isNil(err1)) // 5

	err1 = nil
	fmt.Println(isNil(err1)) // 6
}
