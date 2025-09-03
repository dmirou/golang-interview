// main 1
// child 2
// main 1
// finished
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(1)

	var done atomic.Bool

	fmt.Println("main", runtime.NumGoroutine())

	go func() {
		fmt.Println("child", runtime.NumGoroutine())
		done.Store(true)
	}()

	for !done.Load() {
	}

	fmt.Println("main", runtime.NumGoroutine())

	fmt.Println("finished")
}
