package init

import "fmt"

// init runs automatically
func init() {
	var a [3]int32

	fmt.Printf("a1 %T, len: %d, cap: %d\n", a, len(a), cap(a))
	for i, v := range a {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	var a2 = [1]int32{5}
	fmt.Printf("a2 %T, len: %d, cap: %d\n", a2, len(a2), cap(a2))
	for i, v := range a2 {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	var a3 = [...]int32{1, 2}
	fmt.Printf("a3 %T, len: %d, cap: %d\n", a3, len(a3), cap(a3))
	for i, v := range a3 {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	var a4 = [...]int32{4: 1}
	fmt.Printf("a4 %T, len: %d, cap: %d\n", a4, len(a4), cap(a4))
	for i, v := range a4 {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	var a5 = new([2]int32)
	fmt.Printf("a5 %T, len: %d, cap: %d\n", a5, len(a5), cap(a5))
	for i, v := range a5 {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}
}

func main() {

}
