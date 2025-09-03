package main

import (
	"fmt"
	"slices"
)

func freePortRanges(minPort, maxPort int, busy []int) [][]int {
	busySorted := make([]int, len(busy))
	copy(busySorted, busy)
	slices.Sort(busySorted)

	ranges := make([][]int, 0)

	cur := minPort
	busyIdx := 0

	for {
		if busyIdx == len(busySorted) {
			if cur <= maxPort {
				ranges = append(ranges, []int{cur, maxPort})
			}
			break
		}
		busyVal := busySorted[busyIdx]
		if cur == maxPort {
			if cur != busyVal {
				ranges = append(ranges, []int{cur, cur})
			}
			break
		}

		if cur == busyVal {
			cur++
			busyIdx++
			continue
		}
		if cur > busyVal {
			busyIdx++
			continue
		}
		curMax := min(busyVal-1, maxPort)
		ranges = append(ranges, []int{cur, curMax})
		cur = curMax + 2
		busyIdx++
	}

	return ranges
}

func main() {
	fmt.Println(freePortRanges(1, 3, []int{}))
	fmt.Println(freePortRanges(1, 3, []int{1}))
	fmt.Println(freePortRanges(1, 3, []int{1, 2}))
	fmt.Println(freePortRanges(1, 3, []int{1, 2, 3}))

	fmt.Println(freePortRanges(1, 3, []int{3}))
	fmt.Println(freePortRanges(1, 3, []int{2, 3}))
	fmt.Println(freePortRanges(1, 10, []int{1, 2, 8, 9, 10}))
	fmt.Println(freePortRanges(1, 10, []int{5}))
	fmt.Println(freePortRanges(1, 10, []int{3, 4, 5, 6}))
}
