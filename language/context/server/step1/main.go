package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// slowOperation emulates slow operation which are executing for the specified duration.
// E.g. it can be a database call or some RPC.
func slowOperation(d time.Duration) string {
	log.Println("slowOperation started with duration ", d.String())
	defer log.Println("slowOperation finished with duration ", d.String())
	time.Sleep(d)

	return fmt.Sprintf("result from slowOperation after %s", d.String())
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

		result := slowOperation(d)

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
