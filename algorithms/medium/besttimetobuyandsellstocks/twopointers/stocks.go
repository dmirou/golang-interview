package main

import "fmt"

// 1 2 3
// s f
// s   f
// 2

// 1 2 1
// s f
// 1
// 1 2 1
//   s f

// 5 4 5
// s f
//   s f

func maxProfit(prices []int) int {
	profit := 0

	slow, fast := 0, 1
	last := len(prices) - 1
	for {
		if fast > last {
			break
		}
		if prices[slow] >= prices[fast] {
			slow++
			fast = slow + 1
			continue
		}
		if fast+1 <= last && prices[fast+1] >= prices[fast] {
			fast++
			continue
		}
		profit += prices[fast] - prices[slow]
		slow = fast + 1
		fast = slow + 1
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
