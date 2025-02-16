package main

// 25
// 4 + 25 = 29
// 4 + 81 = 85
// 16 + 25 = 41
// 16 + 1 = 17
// 1 + 49 = 50
// 25 + 0 = 25

// 19
// 1 + 81 = 82
// 64 + 4 = 68
// 36 + 64 = 100
// 1 + 0 + 0 = 1

func isHappy(n int) bool {
	history := make(map[int]struct{})

	for {
		n = sumOfSquares(n)

		if n == 1 {
			return true
		}
		if _, ok := history[n]; ok {
			return false
		}
		history[n] = struct{}{}
	}
}

func sumOfSquares(n int) int {
	sum := 0

	for n > 0 {
		num := n % 10
		sum += num * num
		n = n / 10
	}

	return sum
}

func main() {
	isHappy(19)
}
