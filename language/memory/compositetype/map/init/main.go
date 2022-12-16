package main

import "fmt"

func printDays(name string, m map[int32]string) {
	fmt.Printf("%q => type: %T, ptr: %p, val: %v, len: %d\n", name, m, m, m, len(m))
}

func daysF() {
	var days1 map[int32]string // nil map, len = 0
	printDays("nil days", days1)

	days1 = make(map[int32]string) // not nil map, len = 0
	printDays("not nil days", days1)

	days1[1] = "Monday"
	days1[5] = "Friday"
	printDays("days with values", days1)

	days2 := map[int32]string{
		2: "Tuesday",
		3: "Wednesday",
	}
	printDays("days2", days2)

	fmt.Printf("days2[7] (unexisting): %#v\n", days2[7])

	delete(days2, 2)
	fmt.Printf("days2[2] (deleted): %#v\n", days2[2])
}

type ages map[string]int32

func printAges(name string, m ages) {
	fmt.Printf("%q => type: %T, ptr: %p, val: %v, len: %d\n", name, m, m, m, len(m))
}

func agesF() {
	var a1 = ages{
		"Ivan": 33,
		"Petr": 26,
	}
	printAges("a1", a1)

	a1["NewBorn"]++
	printAges("a1 newBorn", a1)

	if age, ok := a1["unknown"]; !ok {
		fmt.Printf("a1[unknown] = %v", age)
	}
}

func main() {
	fmt.Printf("Days:\n")
	daysF()

	fmt.Printf("\n Ages:\n")
	agesF()
}
