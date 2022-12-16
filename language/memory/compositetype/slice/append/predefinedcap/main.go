package main

import "fmt"

func predefinedCap() {
	data := make([]string, 0, 10)

	lastCap := cap(data)

	for record := 0; record < 10; record++ {

		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			fmt.Printf("predefinedCap record: %d, capacity changed from %d to %d\n", record, lastCap, cap(data))
			lastCap = cap(data)
		}
	}
	fmt.Printf("predefinedCap data: %v\n", data)
}

func predefinedLenAndCap() {
	data := make([]string, 10)

	lastCap := cap(data)

	for record := 0; record < 10; record++ {

		value := fmt.Sprintf("Rec: %d", record)
		data[record] = value

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			fmt.Printf("predefinedLenAndCap record: %d, capacity changed from %d to %d\n", record, lastCap, cap(data))
			lastCap = cap(data)
		}
	}
	fmt.Printf("predefinedLenAndCap data: %v\n", data)
}

func main() {
	predefinedCap()
	predefinedLenAndCap()
}
