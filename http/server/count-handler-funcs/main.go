// Counts amount of request received to any path of http server
// GET /stat returns statistics
// You can send port number to start server on as a first args via cli
// go run main.go 10002
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const defaultPort = 10001

var counts = map[string]int32{}
var countMu = sync.Mutex{}

var total int32

func countHandler(w http.ResponseWriter, r *http.Request) {
	countMu.Lock()
	counts[r.URL.Path]++
	w.Write([]byte(fmt.Sprintf("counts[%s]=%d\n", r.URL.Path, counts[r.URL.Path])))
	total++
	countMu.Unlock()
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	countMu.Lock()
	w.Write([]byte(fmt.Sprintf("total: %d, counts: %v\n", total, counts)))
	countMu.Unlock()
}

func main() {
	port := int64(defaultPort)

	if len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil {
			port = i
		}
	}

	http.HandleFunc("/", countHandler)
	http.HandleFunc("/stat", statHandler)

	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
}
