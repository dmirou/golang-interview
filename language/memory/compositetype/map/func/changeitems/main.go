package main

func foo(m map[int]int) {
	m[10] = 10
	m[15] = 25
}

func main() {
	m := make(map[int]int)
	m[10] = 15
	println("m[10] before foo =", m[10]) // 15
	println("len(m) before foo", len(m)) // 1
	foo(m)
	println("len(m) after foo", len(m)) // 2
	println("m[10] after foo =", m[10]) // 10
	println("m[15] after foo =", m[15]) // 25
}
