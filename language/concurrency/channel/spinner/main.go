package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func fib(x float64) float64 {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spinner(quit <-chan struct{}, done chan<- struct{}) {
	defer fmt.Println("spinner: done")
	for {
		for _, v := range "\\|/*" {
			select {
			case <-quit:
				fmt.Println("spinner: quit received")
				fmt.Println("spinner: complete working")
				close(done)
				return
			default:
				fmt.Printf("\r%c\n", rune(v))
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("can't get input number from args: %v", os.Args)
	}

	number, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("can't parse input number %s: %v", os.Args[1], err)
	}

	quit := make(chan struct{})
	done := make(chan struct{})

	go spinner(quit, done)
	fmt.Println("spinner started")

	result := fib(float64(number))
	close(quit)
	<-done

	fmt.Printf("\rFibonacci(%d) = %.f\n", number, result)
	fmt.Println("main: done")
}
