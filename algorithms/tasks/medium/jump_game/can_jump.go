package jump_game

// Example 1:

//Input: nums = [2,3,1,1,4]
//Output: true
//Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Example 2:
//
//Input: nums = [3,2,1,0,4]
//Output: false
//Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

// dfs
// stack seen
// [0] {}
// if len(stack) > 0
// cur = stack[len(stack) - 1] // 0
// stack = [:len(stack) - 1]
// if seen[cur] skip it
// else
// if cur == len(nums) - 1 { return true }
// if nums[cur] == 0 { continue }
// for i := 1; i <= nums[cur]; i++ {
// stack = append(stack, cur + i)
// }
// seen[cur] = true
// }

//Input: nums = [2,3,1,1,4]
// st: [0] seen: {}
// cur = 0; st: []; seen: {}
// seen[0] != true
// 0 != 4
// st: [1, 2]; seen: {0: true}
//
// cur = 2; st: [1]; seen: {0: true}
// seen[2] != true
// 2 != 4
// st: [1,3], seen: {0: true, 2: true}
//
// cur = 3; st: [1]; seen

func canJump(nums []int) bool {
	st := make([]int, 0, 25)
	st = append(st, 0)
	seen := make(map[int]struct{})

	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		if cur == len(nums)-1 {
			return true
		}
		seen[cur] = struct{}{}
		if nums[cur] == 0 {
			continue
		}
		for i := 1; i <= nums[cur]; i++ {
			possible := cur + i
			if possible > len(nums)-1 {
				continue
			}
			if _, ok := seen[possible]; !ok {
				st = append(st, possible)
			}
		}
	}

	return false
}
