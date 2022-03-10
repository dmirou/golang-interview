package main

import "fmt"

func main() {
	var ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for v := range ch {
		fmt.Printf("%d has been read from the channel\n", v)
	}
	//1 has been read from the channel
	//2 has been read from the channel
	//3 has been read from the channel
}
