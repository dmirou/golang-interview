package main

type Foo struct {
	a int
}

// NewFoo is a good pointer semantics example.
// Guideline:
// Use & with variable in return, not during the struct initialization in factory functions
func NewFoo() *Foo {
	foo := Foo{a: 1}
	// some work with foo here

	return &foo // clear understanding that foo escapes to the heap
}

// NewFoo2 is a good pointer semantics example.
// Guideline:
// Use & with struct to return
func NewFoo2() *Foo {
	return &Foo{a: 1}
}

// NewFoo3 is a bad pointer semantics example.
func NewFoo3() *Foo {
	foo := &Foo{a: 1}
	// some work with foo here

	return foo // don't understand that foo escapes to the heap
}

// NewFoo4 is a terrible pointer / value semantics mix example.
func NewFoo4() Foo {
	foo := &Foo{a: 1}
	// some work with foo here

	return *foo // don't understand that foo escapes to the heap
}

func main() {
}
