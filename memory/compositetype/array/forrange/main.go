// It's copy of https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/arrays/example4/example4.go

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the pointer semantic form of the for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1]) // Betty

	for i := range friends { // uses source array inside the loop
		friends[1] = "Jack" // source array is changed

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1]) // Jack
		}
	}

	// Using the value semantic form of the for range.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends { // make copy of source array and iterate through it
		friends[1] = "Jack" // source array is changed

		if i == 1 {
			// but v has previous value because it's from copy of source array
			fmt.Printf("v[%s]\n", v) // Betty
		}
	}

	fmt.Printf("After loop friends[1] = %v\n", friends[1]) // Jack

	// Using the value semantic form of the for range but with pointer
	// semantic access. DON'T DO THIS.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range &friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
}
