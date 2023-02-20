// Example of HTTP server with graceful shutdown and slow operation with context.
// Handler context can be cancelled only by a client.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// slowOperation emulates slow operation which are executing for the specified duration.
// E.g. it can be a database call or some RPC.
func slowOperation(ctx context.Context, d time.Duration) (string, error) {
	log.Println("slowOperation started with duration ", d.String())
	defer log.Println("slowOperation finished with duration ", d.String())

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(d):
		return fmt.Sprintf("result from slowOperation after %s", d.String()), nil
	}
}

func handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := time.ParseDuration(r.URL.Query().Get("duration"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := io.WriteString(w, fmt.Sprintf("invalid duration: %s\n", d)); err != nil {
				log.Printf("write string: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		result, err := slowOperation(r.Context(), d)
		if err != nil {
			log.Printf("slow operation: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = io.WriteString(w, fmt.Sprintf("%s\n", result))
		if err != nil {
			log.Printf("write string: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(handler()),
	}

	go func() {
		log.Println("server started")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen and serve: %v", err)
			return
		}
		log.Println("listening and serving was stopped")
	}()

	<-ctx.Done()
	log.Println("termination signal received")

	shutdownServer(context.Background(), &srv)
}

func shutdownServer(ctx context.Context, srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v\n", err)
		return
	}
	log.Println("server shutdown completed")
}
