package main

import (
	"fmt"
	"io"
)

type Printer struct {
	io.Writer
}

func (p *Printer) Write(b []byte) (int, error) {
	fmt.Printf("print: %b\n", b)
	return len(b), nil
}

func isNil(name string, i interface{}) {
	if i == nil {
		fmt.Printf("%s (%T) is nil\n", name, i)
		return
	}

	fmt.Printf("%s (%T) is not nil\n", name, i)
}

func main() {
	var w io.Writer
	// w is nil because his dynamic type is nil and his dynamic value is nil
	isNil("w", w)

	var rw io.ReadWriter
	w = rw
	// w is nil because his dynamic type is nil and his dynamic value is nil
	isNil("w", w)

	var p *Printer
	w = p
	// w is not nil because his dynamic type is not nil
	isNil("w", w)
}
