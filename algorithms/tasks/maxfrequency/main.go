package main

import (
	"fmt"
	"sort"
)

/*
Дан массив целых чисел и целое число k. Найти k наиболее часто встречающихся элементов.

Ввод: nums = [3,1,2,1,2,1], k = 2
Вывод: [1, 2]

1. посчитать для каждого эл-та nums сколько раз он встречается (map)
2. отсортировать ключи по убыванию частоты
3. вернуть k первых элементов
*/
func maxFreq(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	keys := make([]int, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	min := k
	if min > len(keys) {
		min = len(keys)
	}

	res := make([]int, min)
	copy(res, keys[:min])

	return res
}

func main() {
	fmt.Printf("%v\n", maxFreq([]int{3, 1, 2, 1, 2, 1}, 2))
	fmt.Printf("%v\n", maxFreq([]int{3, 1, 2, 3, 1, 2, 1, 3, 3}, 2))
}
