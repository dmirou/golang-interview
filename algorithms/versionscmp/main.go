// Нужно разработать метод compare который будет принимать 2 строки, в которой передаётся версия приложения
// Если версия v1 больше чем v2 нужно вернуть 1
// Если версия v1 меньше чем v2 нужно вернуть  -1
// Если версия v1 == v2 нужно вернуть  0
//
// Данные всегда корректны и не может быть букв, не может быть 01, не может отсутствовать цифры
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(v1, v2 string) int {
	items1 := strings.Split(v1, ".")
	items2 := strings.Split(v2, ".")

	n1 := len(items1)
	n2 := len(items2)
	max := n1
	if max < n2 {
		max = n2
	}

	for i := 0; i < max; i++ {
		p1, p2 := 0, 0
		if i < n1 {
			p1, _ = strconv.Atoi(items1[i])
		}
		if i < n2 {
			p2, _ = strconv.Atoi(items2[i])
		}

		switch {
		case p1 < p2:
			return -1
		case p1 > p2:
			return 1
		default:
		}
	}

	return 0
}

func main() {
	fmt.Println(compare("10.0.0", "1.1.1"))  // 1
	fmt.Println(compare("1.1.0.0.1", "1.1")) // 1
	fmt.Println(compare("1.1.0.0", "1.1"))   // 0
	fmt.Println(compare("1.0.0", "1.1.1"))   // -1
	fmt.Println(compare("1.0.0", "1.0.0.1")) // -1

	fmt.Println(compare("1.10.0.0.1", "1.8")) // 1
	fmt.Println(compare("1.0.0", "1.0.0.1"))  // -1
	fmt.Println(compare("1.0.1.1", "1.0.1"))  // 1
	fmt.Println(compare("1.0.1", "1.0.1.1"))  // -1
	fmt.Println(compare("1.81.0", "18.1.0"))  // -1
}
