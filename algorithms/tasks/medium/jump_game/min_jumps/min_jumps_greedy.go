package min_jumps

// [1]
// jumps = 0
// cur = 0
// if cur == n - 1 return

// [2,3,1,1,4]
// jumps = 0
// cur = 0
// if cur == n - 1 return
// jumps = 1
// farthest = cur + nums[0] = 2
// 1 -> 2; step 1
// farthest = max(2, 1+3) = 4
// if farthest >= n-1 return jumps
// farthest = max(4, 2+1) = 4
// if farthest >= n-1 return jumps

// [3, 1, 1, 1, 1, 1]
// jumps = 0
// cur = 0
// farthest = 0
// if cur == n - 1 return
// jumps = 1
// farthest = max(0, 3) = 3
// if farthest >= n-1 return jumps
//

func jump_greedy(nums []int) int {
	jumps := 0
	current := 0
	currentEnd := 0
	farthest := 0

	if current >= len(nums)-1 {
		return jumps
	}

	for {
		if farthest >= len(nums)-1 {
			return jumps + 1
		}
		if current <= currentEnd {
			farthest = max(farthest, current+nums[current])
			current++
			continue
		}
		currentEnd = farthest
		jumps++
	}
}
