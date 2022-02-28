package waitgroup

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const minLatency = 500
const maxLatency = 2000

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	count int32
	mu    sync.Mutex
)

func longHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	ID := count + 1
	count++
	mu.Unlock()
	fmt.Printf("request %d received\n", ID)
	latency := minLatency + rand.Intn(maxLatency-minLatency)
	time.Sleep(time.Duration(latency) * time.Millisecond)
	io.WriteString(w, fmt.Sprintf("%dms. elapsed", latency))

	fmt.Printf("request %d is done after %dms\n", ID, latency)
}

func main() {
	http.HandleFunc("/", longHandler)

	s := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.ListenAndServe(); err != nil {
			log.Printf("server stopped: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		log.Println("shutdown called")
		defer log.Println("shutdown done")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}
	}()

	wg.Wait()
}
