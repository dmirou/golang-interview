package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	type Person struct {
		Name string
	}

	m1 := map[Person]float64{
		Person{Name: "John"}: 4000,
		Person{Name: "Anna"}: 6000,
		Person{Name: "Don"}:  3000,
	}

	keyValueIterator := maps.All(m1)
	fmt.Printf("key value iterator: %#v\n", keyValueIterator)
	keyValueIterator(func(key Person, val float64) bool {
		fmt.Printf("yield true: key: %#v, value: %v\n", key, val)
		return true
	})
	keyValueIterator(func(key Person, val float64) bool {
		fmt.Printf("yield false: key: %#v, value: %v\n", key, val)
		return false
	})
	for k, v := range keyValueIterator {
		p := Person(k)
		fmt.Printf("range: key: %v, value: %v\n", p.Name, v)
	}

	keyIterartor := maps.Keys(m1)
	fmt.Printf("key iterator: %#v\n", keyIterartor)
	keyIterartor(func(p Person) bool {
		fmt.Printf("yield: person: %#v\n", p)
		return true
	})

	for k := range keyIterartor {
		p := Person(k)
		fmt.Printf("range: key: %v\n", p.Name)
	}

	s := []string{"zero", "one", "two", "three"}

	// return only values with even indexes
	evenIterator := func(yield func(val string) bool) {
		for i, line := range s {
			if i%2 == 1 {
				continue
			}

			if !yield(line) {
				return
			}
		}
	}

	for val := range evenIterator {
		fmt.Printf("even iterator: val: %v\n", val)
	}

	slices.All()
}
