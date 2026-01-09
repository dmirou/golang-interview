package rangesum

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	return NumArray{
		nums: nums,
		sums: sums,
	}
}

// nums: [1 2 3 4 5]
// sums: [1 3 6 10 15]
// sumrange (1, 3) = 10-1 = 9
// sumrange (0, 0) = 1-0 = 1
//
// [[[-2,0,3,-5,2,-1]],[2,5]]
// sums [-2 -2 1 -4 -2 -3]
// sumrange (2, 5) = -3 - (-2) = -1
func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	if left == right {
		return this.nums[left]
	}

	return this.sums[right] - this.sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
