package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")
}
