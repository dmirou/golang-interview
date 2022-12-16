// Sample program to show how stacks grow/change.
// goroutine stack is 2 kb by default, see _StackMin = 2048 in runtime/stack.go
// Set stackDebug to 1 in runtime/stack.go to show debug info
package main

// Number of elements to grow each stack frame.
// Run with 1 and then with 1024
const size = 1

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
//go:noinline
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
