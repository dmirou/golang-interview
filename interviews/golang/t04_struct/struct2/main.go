// ./main.go:24:8: ambiguous selector child.Print
// Compilation finished with exit code 1
//
// To fix:
//
//	func (c *Child) Print() {
//		fmt.Println("child")
//	}
package main

import "fmt"

type Parent1 struct{}

func (c *Parent1) Print() {
	fmt.Println("parent1")
}

type Parent2 struct{}

func (c *Parent2) Print() {
	fmt.Println("parent2")
}

type Child struct {
	Parent1
	Parent2
}

func main() {
	var child Child
	child.Print() // ?
}
