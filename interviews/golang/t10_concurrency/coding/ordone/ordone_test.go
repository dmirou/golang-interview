package ordone

import (
	"math/rand"
	"testing"
	"time"
)

func TestOrDone(t *testing.T) {
	n := 1 + rand.Intn(100)
	chs := make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
	}

	done := orDone(chs...)
	go func() {
		i := rand.Intn(n)
		close(chs[i])
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fail()
	}
}
