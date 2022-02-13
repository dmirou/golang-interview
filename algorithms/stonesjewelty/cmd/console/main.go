package main

import (
	"fmt"

	"github.com/dmirou/learngo/algorithms/stonesjewelty/pkg/jewelty"
)

func main() {
	var j, s string

	fmt.Scanf("%s", &j)
	fmt.Scanf("%s", &s)

	c := jewelty.CountFast(string(j), string(s))

	fmt.Printf("%d", c)
}
