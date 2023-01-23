package main

import (
	"fmt"
	"sync/atomic"
)

type Person struct {
	Name string
}

func main() {
	var p atomic.Pointer[Person]

	p.Store(&Person{"Test1"})

	fmt.Printf("%+v\n", p.Load())
}
