package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	x := 1
	y := Ptr(x)

	x = 5
	assert.NotEqual(t, x, *y)
	assert.Equal(t, 1, *y)
}

func TestVal(t *testing.T) {
	assert.Equal(t, nil, Val[any](nil))

	var intNumber int
	var floatNumber float64

	intNumber = 10
	floatNumber = 24.12345

	assert.Equal(t, 10, Val[int](&intNumber))
	assert.Equal(t, 24.12345, Val[float64](&floatNumber))
}

func TestCopy(t *testing.T) {
	type People struct {
		name string
		age  int
	}

	p1 := People{
		name: "Ivan",
		age:  19,
	}

	p2 := Copy(&p1)

	assert.NotEqual(t, &p1, &p2)
	assert.Equal(t, p1.name, p2.name)
	assert.Equal(t, p1.age, p2.age)

	p2.name = "Kristina"
	p2.age = 31

	assert.Equal(t, "Ivan", p1.name)
	assert.Equal(t, 19, p1.age)
}
