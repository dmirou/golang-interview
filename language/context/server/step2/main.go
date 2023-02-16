package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
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
	if err := http.ListenAndServe("localhost:8080", http.HandlerFunc(handler())); err != nil {
		log.Printf("liten and serve: %v", err)
	}
}
