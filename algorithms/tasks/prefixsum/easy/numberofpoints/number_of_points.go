package numberofpoints

/*
1 <= nums.length <= 100
nums[i].length == 2
1 <= starti <= endi <= 100
*/
func numberOfPoints(nums [][]int) int {
	diff := make([]int, 100)

	for _, num := range nums {
		start := num[0] - 1
		diff[start]++

		end := num[1]
		if end == len(diff) {
			continue
		}
		diff[end]--
	}

	current := 0
	points := 0
	for i := 0; i < len(diff); i++ {
		current += diff[i]
		if current > 0 {
			points++
		}
	}

	return points
}
