package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Group struct {
	name    string
	persons []Person
}

func f1(firstName string) {
	p := new(Person) // doesn't escape
	p.firstName = firstName
}

func f2() Group {
	p := &Person{} // doesn't escape

	g := Group{
		persons: []Person{}, // escapes to heap
	}

	g.persons = append(g.persons, *p)

	return g
}

func main() {
	f1("Ivan")
	f2()
}
