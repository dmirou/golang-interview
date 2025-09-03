package main

import "fmt"

type errCustom struct {
	msg string
}

func (e errCustom) Error() string {
	return e.msg
}

func isNil(err error) bool {
	return err == nil
}

// type error interface
//	{ *dynamic_type *data }

func example1() {
	var err1 error           // {*dt = nil, data = nil}
	fmt.Println(isNil(err1)) // true

	var err2 *errCustom      // {*dt = errCustom, data = nil}
	fmt.Println(isNil(err2)) // false

	err2 = &errCustom{"error"} // {*dt = errCustom, data = "error"}
	fmt.Println(isNil(err2))   // false

	err2 = nil               // {*dt = errCustom, data = nil}
	fmt.Println(isNil(err2)) // false

	err1 = err2              // {*dt = nil, data = nil} <- {*dt = errCustom, data = nil}
	fmt.Println(isNil(err1)) // false

	err1 = nil               // {*dt = nil, data = nil}
	fmt.Println(isNil(err1)) // true
}

func main() {
	example1()
}
