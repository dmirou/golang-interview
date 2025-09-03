package main

import "fmt"

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

// is possible too
//func (e MyError) Error() string {
//	return e.msg
//}

func validate() error {
	return &MyError{"This is an error"}
}

func main() {
	fmt.Println(validate())
}
