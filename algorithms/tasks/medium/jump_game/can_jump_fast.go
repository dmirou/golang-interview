package jump_game

// canJumpFast uses a greedy approach to solve the jump game problem.
// Instead of exploring all possible paths (like DFS), it tracks the maximum
// reachable position at each step. This is much more efficient.
func canJumpFast(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	maxReach := 0 // Maximum position we can reach so far

	for i := 0; i < len(nums); i++ {
		// If current position is beyond what we can reach, return false
		if i > maxReach {
			return false
		}

		// Update the maximum reachable position
		// We can reach position i + nums[i] from position i
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		// If we can reach the last index, return true
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return false
}
