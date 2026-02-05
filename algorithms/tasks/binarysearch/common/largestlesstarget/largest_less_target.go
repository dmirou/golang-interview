package largestlesstarget

// Floor / Largest Smaller Than
// Find largest element ≤ target.
// Returns -1 if none exists.
func findFloor(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	floor := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= target {
			floor = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return floor
}
