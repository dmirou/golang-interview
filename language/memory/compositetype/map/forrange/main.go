package main

import "fmt"

func main() {
	var m = map[string]int{}

	m["one"] = 1
	m["two"] = 2

	for key, v := range m {
		fmt.Printf("key: %s, value: %d\n", key, v)
	}

	for key := range m {
		fmt.Printf("key: %s, value: %d\n", key, m[key])
	}
}
