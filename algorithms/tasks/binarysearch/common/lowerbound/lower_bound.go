package lowerbound

/*
First Occurrence / Lower Bound
Find first position where target appears.
*/
func findLowerBound(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	lb := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			lb = mid
			r = mid - 1
			continue
		}
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return lb
}
