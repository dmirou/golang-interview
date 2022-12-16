package main

import (
	"fmt"
	"io"
	"os"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int32
}

func (h Human) String() string {
	return fmt.Sprintf("%s %s, %d years", h.FirstName, h.LastName, h.Age)
}

func show(w interface{}) {
	if w == nil {
		fmt.Printf("%s (%T) is nil\n", "w", w)
	} else if _, ok := w.(io.Closer); ok {
		fmt.Printf("%s (%T) implements %s\n", "w", w, "io.Closer")
	} else {
		fmt.Printf("%s (%T) doesn't implement %s\n", "w", w, "io.Closer")
	}
}

func main() {
	var w io.WriteCloser
	show(w)

	w = os.Stdout
	show(w)

	var s fmt.Stringer
	s = Human{
		FirstName: "Dmitriy",
		LastName:  "Ivanov",
		Age:       25,
	}

	fmt.Printf("s is %q\n", s)

	if h, ok := s.(Human); ok {
		fmt.Printf("s is Human, age is %d\n", h.Age)
	}

	if _, ok := s.(fmt.GoStringer); !ok {
		fmt.Println("s isn't GoStringer")
	}
}
