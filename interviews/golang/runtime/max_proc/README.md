```text
// There is data race.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := false

	fmt.Println("main", runtime.NumGoroutine())

	go func() {
		done = true
		fmt.Println("child", runtime.NumGoroutine())
	}()

	for !done {
	}
	fmt.Println("main", runtime.NumGoroutine())

	fmt.Println("finished")
}
```