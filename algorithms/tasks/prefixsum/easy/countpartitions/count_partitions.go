package countpartitions

func countPartitions(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	partitions := 0
	for i := 0; i < len(sums)-1; i++ {
		rightSum := sums[len(nums)-1] - sums[i]
		if (sums[i]-rightSum)%2 == 0 {
			partitions++
		}
	}
	return partitions
}
