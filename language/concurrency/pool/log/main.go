package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new([1024]int)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log2(w io.Writer, key, val string) {
	b := bufPool.Get().(*[1024]int)
	fmt.Printf("buffer ptr: %d\n", &b[0])
	bufPool.Put(b)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	fmt.Printf("buffer ptr: %d\n", b)
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())
	b.Reset()
	bufPool.Put(b)
}

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	for i := 0; i < 5; i++ {
		fmt.Printf("before log(): alloc = %v bytes\n", mem.Alloc)
		Log2(os.Stdout, "path1", "/search?q=flowers12345")
	}

	// Output: 2006-01-02T15:04:05Z path=/search?q=flowers
}
