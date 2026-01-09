package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dmirou/learngo/algorithms/stonesjewelty/pkg/jewelty"
)

var in, out string

func init() {
	flag.StringVar(&in, "in", "in.txt", "input file")
	flag.StringVar(&out, "out", "out.txt", "output file")
}

func main() {
	flag.Parse()

	in, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	r := bufio.NewReader(in)

	j, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	s, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	c := jewelty.Count(string(j), string(s))

	if err := os.WriteFile(out, []byte(fmt.Sprintf("%d", c)), 0775); err != nil {
		fmt.Println(err)
	}
}
