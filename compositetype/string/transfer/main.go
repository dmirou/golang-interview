package main

import (
	"fmt"
)

func showByValue(name, val string) {
	fmt.Printf(
		"inside showByValue # var: %s, pointer: %p, value: %s, len: %d\n",
		name, &val, val, len(val),
	)
}

func showByPointer(name string, val *string) {
	fmt.Printf(
		"inside showByPointer # var: %s, pointer: %p, value: %s, len: %d\n",
		name, val, *val, len(*val),
	)
}

func main() {
	var s = "hello"
	fmt.Printf(
		"outside showByPointer # var: %s, pointer: %p, value: %s, len: %d\n",
		"s", &s, s, len(s),
	)
	showByPointer("s", &s)
	showByValue("s", s)

	s = "привет"
	fmt.Printf(
		"outside showByPointer # var: %s, pointer: %p, value: %s, len: %d\n",
		"s", &s, s, len(s),
	)
	showByPointer("s", &s)
	showByValue("s", s)

	s2 := s
	showByPointer("s2", &s2)
}
