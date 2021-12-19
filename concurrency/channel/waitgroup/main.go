package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WaitGroup struct {
	sema  chan struct{}
	count int
}

func NewWg() *WaitGroup {
	return &WaitGroup{
		sema: make(chan struct{}, 1),
	}
}

func (wg *WaitGroup) Inc() {
	wg.sema <- struct{}{}
	wg.count++
	<-wg.sema
}

func (wg *WaitGroup) Done() {
	wg.sema <- struct{}{}
	wg.count--
	<-wg.sema
}

func (wg *WaitGroup) Wait() {
	done := false
	for !done {
		wg.sema <- struct{}{}
		if wg.count == 0 {
			done = true
		}
		<-wg.sema
	}
}

func doLongWork(idx int, out chan string) {
	rand.Seed(time.Now().UnixNano())
	wait := time.Duration(rand.Int31n(3000))
	time.Sleep(time.Duration(wait) * time.Millisecond)

	out <- fmt.Sprintf("%d done after %d", idx, wait)
}

func main() {
	var wg = NewWg()
	out := make(chan string)

	go func(ch chan string) {
		for msg := range ch {
			fmt.Println(msg)
		}
		fmt.Println("print done")
	}(out)

	for i := 1; i < 20; i++ {
		wg.Inc()
		go func(idx int, out chan string, wg *WaitGroup) {
			doLongWork(idx, out)
			wg.Done()
		}(i, out, wg)
	}

	wg.Wait()
	close(out)
}
