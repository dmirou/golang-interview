// A type assertion provides access to an interface value's
// underlying concrete value.
// t := i.(T)
// This statement asserts that the interface value i holds the concrete type T
// and assigns the underlying T value to the variable t.
// If i does not hold a T, the statement will trigger a panic.
//
// To test whether an interface value holds a specific type, a type assertion
// can return two values: the underlying value and a boolean value that reports
// whether the assertion succeeded.
// t, ok := i.(T)
// If i holds a T, then t will be the underlying value and ok will be true.
// If not, ok will be false and t will be the zero value of type T, and no panic occurs.
package main

import (
	"fmt"
	"io"
	"os"
)

type Connection struct {
	io.ReadWriteCloser
}

func checkType(rwc io.ReadWriteCloser) {
	switch val := rwc.(type) {
	case io.Reader:
		fmt.Printf("reader dynamic value: %#v, dynamic type: %T\n", val, val)
	default:
		fmt.Printf("unexpected value: %#v, type: %T\n", val, val)
	}
}

func checkAnyType(i any) {
	switch val := i.(type) {
	case io.Reader:
		fmt.Printf("reader dynamic value: %#v, dynamic type: %T\n", val, val)
	case io.Writer:
		fmt.Printf("writer dynamic value: %#v, dynamic type: %T\n", val, val)
	default:
		fmt.Printf("unexpected value: %#v, type: %T\n", val, val)
	}
}

func main() {
	var c io.ReadWriteCloser = Connection{}
	checkType(c)
	// reader dynamic value: main.Connection{ReadWriteCloser:io.ReadWriteCloser(nil)}, dynamic type: main.Connection

	var i io.ReadWriteCloser
	checkType(i)
	//unexpected value: <nil>, type: <nil>

	var f io.ReadWriteCloser = os.Stdout
	checkType(f)
	//reader dynamic value: &os.File{file:(*os.file)(0xc0001180c0)}, dynamic type: *os.File

	type w1 struct {
		io.Writer
	}
	var w w1
	checkAnyType(w)
	//writer dynamic value: main.w1{Writer:io.Writer(nil)}, dynamic type: main.w1
}
