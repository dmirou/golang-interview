package main

// changes stayed only in local variable (pointer to hmap)
// there will be no changes in received map
func foo(m map[int]int) {
	m = make(map[int]int)
	m[10] = 12
	m[43] = 13
}

func main() {
	m := make(map[int]int)
	m[10] = 15
	println("m[10] before foo =", m[10]) // 15
	println("len(m) before foo", len(m)) // 1
	foo(m)
	println("len(m) after foo", len(m)) // 1
	println("m[10] after foo =", m[10]) // 15
}
