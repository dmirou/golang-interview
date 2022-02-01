package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type pubApi struct{}

func (pubApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from public api")
}

func NewPublicServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Hello from /\n")
	})

	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/test/ request from %s", r.URL.String())
	})

	mux.Handle("/api/", pubApi{})

	return &http.Server{
		Addr:    "localhost:8090",
		Handler: mux,
	}
}

func NewPrivateServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "homepage of %s\n", r.URL.String())
	})

	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/test/ request from %s", r.URL.String())
	})

	return &http.Server{
		Addr:    "localhost:8091",
		Handler: mux,
	}
}

func main() {
	//var done chan struct{}

	publicSrv := NewPublicServer()

	go func() {
		log.Printf("start public server at %v", publicSrv.Addr)
		err := publicSrv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Printf("unexpected server error: %v", err)
		}
		log.Printf("server %s closed", publicSrv.Addr)
	}()

	privateSrv := NewPrivateServer()

	go func() {
		log.Printf("start private server at %v", privateSrv.Addr)
		err := privateSrv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Printf("unexpected server error: %v", err)
		}
		log.Printf("server %s closed", privateSrv.Addr)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	err := publicSrv.Shutdown(ctx)
	if err != nil {
		log.Printf("server %s shutdown complete with error: %v", publicSrv.Addr, err)
	} else {
		log.Printf("server %s shutdown complete", publicSrv.Addr)
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	err = privateSrv.Shutdown(ctx)
	if err != nil {
		log.Printf("server %s shutdown complete with error: %v", privateSrv.Addr, err)
	} else {
		log.Printf("server %s shutdown complete", privateSrv.Addr)
	}
}
