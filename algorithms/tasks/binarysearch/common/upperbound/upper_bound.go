package upperbound

// Last Occurrence / Upper Bound
// Find last position where target appears.
// Returns -1 if target wasn't find.
func findUpperBound(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	ub := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			ub = mid
			l = mid + 1
			continue
		}
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ub
}
