package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(i)
		}()
	}
}
