// panic: strings: illegal use of non-zero Builder copied by value
// to fix: remove copyBuilder func and
// sb2 := strings.Builder{}
// sb2.WriteString(sb.String())
// output:
// hello
// helloworld
package main

import (
	"fmt"
	"strings"
)

func copyBuilder(sb strings.Builder) strings.Builder {
	return sb
}

func main() {
	sb := strings.Builder{}
	sb.Grow(10)
	sb.WriteString("hello")

	sb2 := copyBuilder(sb)
	sb2.WriteString("world")

	fmt.Println(sb.String())  // ?
	fmt.Println(sb2.String()) // ?
}
