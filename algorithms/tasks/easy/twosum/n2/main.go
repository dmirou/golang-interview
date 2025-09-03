// https://leetcode.com/problems/two-sum/
package main

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
// Input: nums = [3,3], target = 6
//Output: [0,1]
//
// Constraints:
//
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
//
// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
func twoSumN2(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {

}
