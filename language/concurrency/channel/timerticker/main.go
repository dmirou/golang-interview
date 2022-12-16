package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)

cycle:
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Printf("%v tick\n", time.Now().Format(time.Stamp))
		case <-timer.C:
			fmt.Printf("%v timer stopped\n", time.Now().Format(time.Stamp))
			break cycle
		}
	}
}
