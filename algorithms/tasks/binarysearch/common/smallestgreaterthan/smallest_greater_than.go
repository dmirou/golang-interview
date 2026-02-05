package smallestgreaterthan

// Ceil / Smallest Greater Than
// Find smallest element ≥ target.
// Returns -1 if none exists.
func findCeil(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	ceil := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] >= target {
			ceil = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ceil
}
