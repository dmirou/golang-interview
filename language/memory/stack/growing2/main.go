package main

func main() {
	var x [10]int
	println(`main x`, &x)
	a(x)
	println(`main x`, &x)
}

//go:noinline
func a(x [10]int) {
	println(`func a`)
	var y [100]int
	println(`a y`, &y)
	b(y)
	println(`a y`, &y)
}

//go:noinline
func b(x [100]int) {
	println(`func b`)
	var y [1000]int
	println(`b y`, &y)
	c(y)
	println(`b y`, &y)
}

//go:noinline
func c(x [1000]int) {
	println(`func c`)
}
