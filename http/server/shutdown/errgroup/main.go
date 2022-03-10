package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
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
	sm := http.NewServeMux()
	sm.HandleFunc("/", longHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: sm,
	}

	rootCtx := context.Background()
	g, ctx := errgroup.WithContext(rootCtx)

	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)
		defer stop()

		<-ctx.Done()
		err := srv.Shutdown(context.Background())
		if err != nil {
			fmt.Printf("server shutdown finished with error: %v\n", err)
		} else {
			fmt.Printf("server shutdown finished\n")
		}

		return err
	})

	err := g.Wait()
	fmt.Printf("main done, err: %v\n", err)
}
