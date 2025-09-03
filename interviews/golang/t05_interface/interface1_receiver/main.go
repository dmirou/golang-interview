// MyFloat implements Abser
// *Vertex implements Abser
// error: cannot use v (variable of struct type Vertex) as Abser value in assignment:
// Vertex does not implement Abser (method Abs has pointer receiver)
// To fix: printAbs(&v)
// 1.25
// 5
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func printAbs(a Abser) {
	fmt.Println(a.Abs())
}

func main() {
	f := MyFloat(-1.25)
	v := Vertex{3, 4}

	printAbs(f) // 1
	printAbs(v) // 2
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
