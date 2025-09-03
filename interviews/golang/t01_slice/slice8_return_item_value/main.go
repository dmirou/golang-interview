package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("before get(): alloc = %v KiB\n", mem.Alloc/1024)

	resOfRes := make([][]int, 0, 100)

	for i := 0; i < 50; i++ {
		res := getLastElem()

		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("after get(): alloc = %v KiB, slice: %v\n", mem.Alloc/1024, res)

		resOfRes = append(resOfRes, res)
	}

	fmt.Println()
	resOfRes = nil

	for i := 0; i < 20000; i++ {
		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("alloc = %v KiB\n", mem.Alloc/1024)
		time.Sleep(time.Second * 2)
	}

	fmt.Println(resOfRes)
}

func getLastElem() []int {
	sl := make([]int, 0, 100_000)
	for i := 0; i < cap(sl); i++ {
		sl = append(sl, i)
	}

	res := sl[99_999:]
	cpy := make([]int, len(res))
	copy(cpy, res)

	return cpy
}
