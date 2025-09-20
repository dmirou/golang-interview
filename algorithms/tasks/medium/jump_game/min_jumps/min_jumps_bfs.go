package min_jumps

// [2,3,1,1,4]
// q {0} j = 0
// q {2, 1} j = 1
// q {2, 1, 3, 4}, 4 == n - 1, return j+1

// q {0} seen {0}  jumps = 0
// q {1, 2} seen {0} jumps = 1
// ql = 2, cur = 1, q{2, 4, 3, 2}, jumps = 2
// ql = 2, cur = 2, q{4, 3, 2, 3}, jumps = 2
// if cur == n - 1 return jumps

// [2,3,0,1,4]
// q {0}, seen{} j = 0
// ql = 1, cur = 0, nums[0] == 2, q{2, 1}, jumps = 1
// ql = 2, cur = 2, nums[2] = 0, q{1}, jumps = 2
// ql =2, cur = 1, nums[1] = 3, q{4, 3, 2}, jumps = 2
// ql = 3, cur = 4, cur == n - 1, return jumps

// [1]
// q{0} seen {}, j = 0
// ql = 1, cur = 0, 0 == 1 -1 return j

// [3, 1, 1, 1, 1, 1]
// q {0}
// q {0, 3, 2, 1}
// q {3, 2, 1}
// q {3, 2, 1, 4}
// q {3, 2, 1, 4, 3}

func jump_bfs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	seen := make(map[int]struct{})
	jumps := 0
	queue := make([]int, 1)
	queue[0] = 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			start := queue[i]
			seen[start] = struct{}{}
			for j := nums[start]; j > 0; j-- {
				if start+j >= len(nums)-1 {
					return jumps + 1
				}
				if _, ok := seen[start+j]; ok {
					continue
				}
				queue = append(queue, start+j)
			}
		}
		queue = queue[levelSize:]
		jumps++
	}

	return jumps
}
