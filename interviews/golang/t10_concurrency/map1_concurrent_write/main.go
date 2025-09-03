// fatal error: concurrent map writes
package main

import (
	"fmt"
	"time"
)

func main() {
	var data = make(map[int]bool)
	for i := 0; i < 10; i++ {
		go func() {
			data[i] = true
		}()
	}

	time.Sleep(time.Second)
	for k, v := range data {
		fmt.Println(k, v) // ?
	}
}
