package main

import (
	"fmt"
	"time"
)

func d0() (x int) {
	x = 1

	defer fmt.Printf("1. defer fmt.Printf(x) is %d\n", x)

	x++

	defer fmt.Printf("2. defer fmt.Printf(x) is %d\n", x)

	defer func() {
		fmt.Printf("3. defer func(){fmt.Printf(x)} is %d\n", x)
	}()

	defer func(x int) {
		x += 2
		fmt.Printf("4. x increased in defer func(x int) {x+=2} from %d to %d\n", x-2, x)
	}(x)

	defer func(x *int) {
		*x++
		fmt.Printf("5. x increased in defer func(x *int) {*x++} from %d to %d\n", *x-1, *x)
	}(&x)

	defer func() {
		x++
		fmt.Printf("6. x increased in defer func() {x++} from %d to %d\n", x-1, x)
	}()

	return x
}

func d1() {
	x := 1

	defer func() {
		fmt.Printf("defer1 func(): x = %d\n", x)
	}()

	defer func(x int) {
		fmt.Printf("defer2 func(x int): x = %d\n", x)
	}(x)

	x = 2
}

func d2() int {
	x := 1

	defer func(x *int) {
		*x = 5
		fmt.Printf("defer2 func(x *int) *x = 5\n")
	}(&x)

	fmt.Printf("x = %d\n", x) // x = 1

	return x
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(100 * time.Millisecond)
}

func double() (result int) {
	defer func() {
		v := recover()
		fmt.Printf("recover called, v is %q\n", v)
		result = 10
		fmt.Println("defer called result = 10")
	}()

	result = 5
	panic("some panic reason")
}

func trace(name string) func() {
	start := time.Now()
	fmt.Printf("%s started at %s\n", name, start)

	return func() {
		fmt.Printf("%s finished at %s after %s\n", name, time.Now(), time.Since(start))
	}
}

func d3() {
	fmt.Println("d3")
	defer fmt.Println()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	return
}

func d4() {
	fmt.Println("d4")
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}

	return
}

func d5() {
	fmt.Println("d5")
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Printf("%d ", i)
		}(i)
	}

	return
}

func d6() {
	fmt.Println("d6")
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}

	return
}

func main() {
	fmt.Println("d0")
	x := d0()
	fmt.Printf("result from d0 is %d\n\n", x)

	fmt.Println("d1")
	d1()
	// defer2 func(x int): x = 1
	// defer1 func(): x = 2

	fmt.Printf("d2 is %d\n\n", d2())
	// x = 1
	// defer2 func(x *int) *x = 5
	// d2 is 1

	fmt.Printf("double() is %d\n\n", double())
	// double() is 10

	bigSlowOperation()
	// bigSlowOperation started at 2022-02-10 20:29:46.345536345 +0700 +07 m=+0.000268892
	// bigSlowOperation finished at 2022-02-10 20:29:48.347305095 +0700 +07 m=+2.002037770 after 2.001771322s

	d3()
	// 4 3 2 1 0

	d4()
	// 5 5 5 5 5

	d5()
	// 4 3 2 1 0

	d6()
	// 4 3 2 1 0
}
