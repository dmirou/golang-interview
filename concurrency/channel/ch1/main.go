// Syncs main and child goroutines by unbuffered channel.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		for i := 1; i < 20; i++ {
			fmt.Printf("%v, routine\n", time.Now().Format("15:04:05"))
			time.Sleep(100 * time.Millisecond)
		}

		ch <- struct{}{}
	}()

	<-ch

	fmt.Printf("%v, main done\n", time.Now().Format("15:04:05"))
}
