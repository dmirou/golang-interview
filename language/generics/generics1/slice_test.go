package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSum(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 6, SliceSum([]int{3, 3, 2, -2}))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, 6.7, SliceSum([]float64{3.5, 3.5, 2.0, -2.3}))
	})
}

func TestSliceApplyFuncInMemory(t *testing.T) {
	type Person struct {
		firstName  string
		secondName string
		fullName   string
	}

	people := []Person{
		{
			firstName:  "Ivan",
			secondName: "Ivanov",
		},
		{
			firstName:  "Petr",
			secondName: "Petrov",
		},
	}

	SliceApplyFuncInMemory(people, func(src *Person) {
		src.fullName = fmt.Sprint(src.firstName, " ", src.secondName)
	})

	assert.Equal(t, "Ivan Ivanov", people[0].fullName)
	assert.Equal(t, "Petr Petrov", people[1].fullName)
}

func TestSliceApplyFunc(t *testing.T) {
	type Person struct {
		firstName  string
		secondName string
		fullName   string
	}

	people := []Person{
		{
			firstName:  "Ivan",
			secondName: "Ivanov",
		},
		{
			firstName:  "Petr",
			secondName: "Petrov",
		},
	}

	results := SliceApplyFunc(people, func(p Person) Person {
		p.fullName = fmt.Sprint(p.firstName, " ", p.secondName)
		return p
	})

	assert.Equal(t, "", people[0].fullName)
	assert.Equal(t, "", people[1].fullName)
	assert.Equal(t, "Ivan Ivanov", results[0].fullName)
	assert.Equal(t, "Petr Petrov", results[1].fullName)
}
