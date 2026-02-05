package main

import "fmt"

// На вход подаются два неупорядоченных слайса любой длины.
// Надо написать функцию, которая возвращает их пересечение
// [], [] -> []
// [1], [1] -> [1]
// [1, 1], [1, 2] -> [1]
// [1, 2], [1, 1] -> [1]
func intersectionN2(a, b []int) []int {
	isn := make([]int, 0)
	used := make(map[int]bool)

	for _, va := range a {
		for j, vb := range b {
			if va == vb && !used[j] {
				isn = append(isn, va)
				used[j] = true
				break
			}
		}
	}

	return isn
}

func intersectionFast(a, b []int) []int {
	isn := make([]int, 0)
	counts := make(map[int]int)

	min := a
	max := b
	if len(min) > len(b) {
		min = b
		max = a
	}
	for _, v := range min {
		counts[v]++
	}

	for _, v := range max {
		if counts[v] > 0 {
			counts[v]--
			isn = append(isn, v)
		}
	}

	return isn
}

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("intersectionN2 %v\n", intersectionN2(a, b))
	fmt.Printf("intersectionFast %v\n", intersectionFast(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("intersectionN2(a,b) %v\n", intersectionN2(a, b))
	fmt.Printf("intersectionFast(a,b) %v\n", intersectionFast(a, b))
	fmt.Printf("intersectionN2(b,a) %v\n", intersectionN2(b, a))
	fmt.Printf("intersectionFast(b,a) %v\n", intersectionFast(b, a))
	a = []int{1, 2, 2, 3, 4, 1}
	b = []int{1, 1, 2, 5, 6, 2}
	// [1, 2, 2, 1]
	fmt.Printf("intersectionN2(a,b) %v\n", intersectionN2(a, b))
	fmt.Printf("intersectionFast(a,b) %v\n", intersectionFast(a, b))
	fmt.Printf("intersectionN2(b,a) %v\n", intersectionN2(b, a))
	fmt.Printf("intersectionFast(b,a) %v\n", intersectionFast(b, a))
}
