// &{Ivan 20}
// &{Ivan 20}
// To fix:
//   - option 1: in changePerson: *p = Person
//   - option 2:
//     func changePerson(p **Person)
//     in: *p = &Person
//     call: changePerson(&person)
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changePerson(p *Person) {
	p = &Person{
		name: "Vladimir",
		age:  22,
	}
}

func main() {
	person := &Person{
		name: "Ivan",
		age:  20,
	}
	fmt.Println(person) // 1

	changePerson(person)
	fmt.Println(person) // 2
}
