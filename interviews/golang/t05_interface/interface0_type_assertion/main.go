package main

import "fmt"

type Foo struct{}

func (a *Foo) A() {}
func (a *Foo) B() {}
func (a *Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}
	fmt.Printf("%T\n", f)

	y := f.(BC)
	fmt.Printf("%T\n", y)

	y.C()
	// y.A()
}
