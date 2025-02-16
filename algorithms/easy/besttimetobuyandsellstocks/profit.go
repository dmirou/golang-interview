package besttimetobuyandsellstocks

import "math"

// 7 2 8 1 3
// buy = 7
// if buy > prices[i] buy = prices[i], buy = 2
// if buy < prices[i] && sell == -1, profit = prices[i] - buy = 8 - 2 = 5
// if buy > prices[i], buy = prices[i], buy = 1
// if buy < prices[i], profit = max(profit, prices[i] - buy)
func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0.0

	for i := 1; i < len(prices); i++ {
		switch {
		case buy > prices[i]:
			buy = prices[i]
		case buy < prices[i]:
			profit = math.Max(profit, float64(prices[i]-buy))
		default:
			continue
		}
	}

	return int(profit)
}
