package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
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
	start := time.Now()
	mu.Lock()
	ID := count + 1
	count++
	mu.Unlock()
	fmt.Printf("request %d received\n", ID)
	latency := minLatency + rand.Intn(maxLatency-minLatency)

	for {
		select {
		case <-r.Context().Done():
			fmt.Printf("long running request %d done after %v\n", ID, time.Since(start))
			fmt.Println("request graceful handler exit")
			return
		case <-time.After(time.Duration(latency) * time.Millisecond):
			io.WriteString(w, fmt.Sprintf("long running request %d sent new data chunk\n", ID))
		}
	}
}

func main() {
	g, mainCtx := errgroup.WithContext(context.Background())

	sm := http.NewServeMux()
	sm.HandleFunc("/", longHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: sm,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		},
	}

	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		ctx, stop := signal.NotifyContext(mainCtx, syscall.SIGTERM, os.Interrupt)
		defer stop()

		<-ctx.Done()

		shctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(shctx)
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
