package main

import (
	"time"
)

func main() {
	{
		ch := make(chan int, 3) // qcount 0, dataqsiz 3, sendx 0, recvx 0

		ch <- 1 // qcount 1, dataqsiz 3, buf [...] (len:3), sendx 1, recvx 0
		ch <- 2 // qcount 2, dataqsiz 3, buf [...] (len:3), sendx 2, recvx 0

		<-ch      // qcount 1, dataqsiz 3, buf [...] (len:3), sendx 2, recvx 1
		close(ch) // closed 1
	}

	{
		ch2 := make(chan int) // qcount 0, dataqsiz = 2, buf [...] (len:2)

		go func() {
			for {
				<-ch2
			}
		}()

		time.Sleep(time.Second)
		time.Sleep(time.Second)
	}
}
