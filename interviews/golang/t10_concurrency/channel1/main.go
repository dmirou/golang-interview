// program freezes because channel is not initialized
// to fix:
// ch = make(chan string)
// move ch <- "world" after go func
// or
// ch = make(chan string, 1)
// after that:
// hello world; main done; child done
// or hello world; child done; main done
package main

import "fmt"

func main() {
	defer fmt.Println("main done")

	var ch chan string

	ch <- "world"

	go func() {
		defer fmt.Println("child done")
		fmt.Println("hello", <-ch)
	}()

}
