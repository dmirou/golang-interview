package main

import "fmt"

/*
Нужно слить 2 отсортированных массива в один отсортированный массив

Пример
Ввод [1, 2, 5], [1, 2, 3, 4, 6]
Вывод [1, 1, 2, 2, 3, 4, 5, 6]

Ввод [1, 1, 2], [1, 2, 3]
Вывод [1, 1, 1, 2, 2, 3]

[1] [1] => [1, 1]
[2] [1] => [1, 2]
[2] [1, 1] => [1, 1, 2]

[1] [2,3]
*/
func merge(a, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}

	var res []int

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	if a[0] < b[0] {
		res := append(res, a[0])
		return append(res, merge(a[1:], b)...)
	}

	res = append(res, b[0])
	return append(res, merge(a, b[1:])...)
}

func main() {
	fmt.Printf("%v\n", merge(nil, nil))
	fmt.Printf("%v\n", merge([]int{1}, []int{2, 3}))
	fmt.Printf("%v\n", merge([]int{1, 1, 2, 3, 10}, []int{1, 2, 2, 4, 11}))
}
