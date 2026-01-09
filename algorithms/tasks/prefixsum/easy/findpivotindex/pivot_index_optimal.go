package findpivotindex

// Input: nums = [1,7,3,6,5,6]
// Output: 3
// ts = 28
// i 0, ls 0, rs = 28-1 = 27
// i 1, ls = 0+1 = 1, rs = rs-7= 27-7=20
// i 2, ls = 1 +7 = 8, rs = 20-3=17
// i 3, ls =8+3 = 11, rs = 17-6 = 11

// Input: nums = [1,2,3]
// Output: -1
// ts = 6
// i 0, ls 0, rs = 6-1-0 = 5
// i 1, ls 1, rs = 5-2 = 3
// i 2, ls 3, rs = 3-3 = 0

// Input: nums = [2,1,-1]
// Output: 0
// ts = 2
// i = 0, ls = 0, rs = ts - ls - nums[0] = 0,
// if ls == ts, return i

// Input: nums = [0]
// Output: 0

func pivotIndexOptimal(nums []int) int {
	ts := 0
	for _, v := range nums {
		ts += v
	}
	ls := 0

	for i := 0; i < len(nums); i++ {
		rs := ts - ls - nums[i]
		if ls == rs {
			return i
		}
		ls += nums[i]
	}
	return -1
}
