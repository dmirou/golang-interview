package stack

/*
813. Largest Sum of Averages

Medium
Topics
premium lock icon
Companies

You are given an integer array nums and an integer k. You can partition the array
into at most k non-empty adjacent subarrays. The score of a partition is the sum
of the averages of each subarray.
Note that the partition must use every integer in nums, and that the score is not
necessarily an integer.
Return the maximum score you can achieve of all the possible partitions.
Answers within 10-6 of the actual answer will be accepted.

Example 1:
Input: nums = [9,1,2,3,9], k = 3
Output: 20.00000
Explanation:
The best choice is to partition nums into [9], [1, 2, 3], [9].
The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
We could have also partitioned nums into [9, 1], [2], [3, 9], for example.
That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.

Example 2:
Input: nums = [1,2,3,4,5,6,7], k = 4
Output: 20.50000

Constraints:
1 <= nums.length <= 100
1 <= nums[i] <= 104
1 <= k <= nums.length
========0   1   2   3   4  5
nums = [9,  1,  2,  3,  9]
sums = [0,  9, 10, 12, 15, 24]
sum[1, 3] = sums[4] - sums[1] = 15 - 9 = 6
sum[2, 4] = sums[5] - sums[2] = 24 - 10 = 14
sum[0, 0] = sums[1] - sums[0] = 9 - 0 = 0
sum[i, j] = sums[j+1] - sums[i]
*/
func largestSumOfAverages(nums []int, k int) float64 {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	return largestAvgSum(sums, 0, len(nums)-1, k)
}

/*
sum[i, j] = sums[j+1] - sums[i]

[9],[1],[2,3,9]
[9],[1,2],[3,9]
[9],[1,2,3],[9]
[9],[1,2,3,9]
[9,1],[2],[3,9]
[9,1],[2,3],[9]
[9,1],[2,3,9]
[9,1,2],[3],[9]
[9,1,2,3],[9]
[9,1,2,3,9]

////////0 1  2  3  4  5
nums = [9 1  2  3  9], k = 3
sums = [0 9 10 12 15 24]

largestAvgSum, start0 end4 k3
start0 > end4 false
k3==1 false || start0==end4 false
result = 0
i0 i<end4
result = max(0, avgsum(0,0) + largestAvgSum(1, 4, 2))

largestAvgSum, start1, end4, k2
result = 0
i1<end4
result = max(0, avgsum(1,1) + largestAvgSum(2, 4, 1)) = max(0, 1 + 4.66) = 5.66
i2<end4
result = max(5.66, avgsum(1,2) + largestAvgSum(3, 4, 1)) = max(5.66, 1.5 + 6) = 7.5

largestAvgSum, start2, end4, k1
k3==1 true, return avgSum(2, 4) = (24 - 10) / (4-2+1) = 14 / 3 = 4.66
*/
/*
largestSumOfAverages([1 2 3 4 5 6 7], 4) = 20.000000, want 20.500000

largestSumOfAverages([1 2 3 4 5 6 7], 4)
               sums[0 1 3 6 10 15 21 28]
largestAvgSum([0 1 3 6 10 15 21 28], 0, 6, 4)
max(0, 1 + largestAvgSum(sums, 1, 6, 3))
max(0, 2 + largestAvgSum(sums, 2, 6, 2))
*/
type Frame struct {
	start, end int
	k          int
	i          int // current split position
	localMax   float64
	phase      int // 0: initial, 1: processing splits
}

func largestAvgSum(sums []int, start, end int, k int) float64 {
	if start > end {
		return 0
	}
	if k == 1 || start == end {
		return avgSum(sums, start, end)
	}

	stack := []Frame{{start, end, k, start, 0, 0}}
	result := 0.0

	for len(stack) > 0 {
		frame := &stack[len(stack)-1]

		// Base cases
		if frame.start > frame.end {
			result = 0
			stack = stack[:len(stack)-1]
			continue
		}
		if frame.k == 1 || frame.start == frame.end {
			result = avgSum(sums, frame.start, frame.end)
			stack = stack[:len(stack)-1]
			continue
		}

		if frame.phase == 0 {
			// Initialize processing
			frame.localMax = 0
			frame.i = frame.start
			frame.phase = 1
		}

		if frame.phase == 1 {
			if frame.i < frame.end {
				// Need to compute recursive call for right side
				// Push frame to compute right side
				stack = append(stack, Frame{
					start:    frame.i + 1,
					end:      frame.end,
					k:        frame.k - 1,
					i:        frame.i + 1,
					localMax: 0,
					phase:    0,
				})
				frame.phase = 2 // Will resume after recursive call
				frame.i++       // Move to next i for later
			} else {
				// Done with all splits
				result = frame.localMax
				stack = stack[:len(stack)-1]
			}
		} else if frame.phase == 2 {
			// Returning from recursive call
			current := avgSum(sums, frame.start, frame.i-1) + result
			frame.localMax = max(frame.localMax, current)
			frame.phase = 1 // Continue with next split
		}
	}

	return result
}

func avgSum(sums []int, start, end int) float64 {
	return float64(sums[end+1]-sums[start]) / float64(end-start+1)
}
