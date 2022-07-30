package main

func f1(a, b int) int {
	c := a + b
	d := a - b

	return c - d
}

func f2() *int {
	i := 10 // moved to heap
	return &i
}

func f31() string {
	m := make(map[int]string) // doesn't escape
	m[6] = "test"

	return m[6]
}

func f32() map[int]string {
	m := make(map[int]string) // escapes to heap
	return m
}

func f41() string {
	x := make([]string, 1) // doesn't escape
	x[0] = "ab"
	return x[0]
}

func f42() []string {
	x := make([]string, 0) // escapes to heap
	return x
}

func f5(a, b []string) []string {
	c := append(a, b...)
	return c
}

var g1 = 5

func f61() int {
	g1++

	if g1 > 0 {
		return 10
	}
	return g1
}

func f62() int {
	g1++

	if g1 > 10 {
		return 25
	}
	return g1
}

func main() {
	f1(1, 10)
	f2()

	f31()
	f32()

	f41()
	f42()

	a := []string{"test", "test2"}  // doesn't escape
	b := []string{"test4", "test5"} // doesn't escape
	f5(a, b)

	f61()
	f62()
}
