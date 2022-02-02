package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	var nc net.Conn
	var rwc io.ReadWriteCloser
	var rw io.ReadWriter
	var r io.Reader

	rwc = nc
	rw = rwc
	rw = &bytes.Buffer{}
	io.WriteString(rw, "test")
	fmt.Printf("rw (%T) is %v\n", rw, rw)
	r = rw
	fmt.Printf("r (%T) is %v\n", r, r)
}
