package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	ranges := make([]string, 0)
	starti := 0
	r := ""
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		if starti == i-1 {
			r = fmt.Sprint(nums[starti])
		} else {
			r = fmt.Sprintf("%d->%d", nums[starti], nums[i-1])
		}
		ranges = append(ranges, r)
		starti = i
	}

	if starti == len(nums)-1 {
		r = fmt.Sprint(nums[starti])
	} else {
		r = fmt.Sprintf("%d->%d", nums[starti], nums[len(nums)-1])
	}
	ranges = append(ranges, r)

	return ranges
}
