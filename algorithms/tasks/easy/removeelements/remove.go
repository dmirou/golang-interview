package removeelements

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are
// not equal to val.
// Consider the number of elements in nums which are not equal to val be k, to get accepted,
// you need to do the following things:
// Change the array nums such that the first k elements of nums contain the elements which are not
// equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.

// 1 2 3 2 1 # 2
// i=0, left = 0, right = 4
// 1 != 2, i++, i = 1, left++, left =1
// 2 == 2, i++, i = 2, left=1
// 3 == 2, left < i, nums[left] = 3, left++, left =2, i++, i = 3
// 2 == 2, i++, i =4
// 1 != 2, left < i, numbs[left] = 1, i++

func removeElement(nums []int, val int) int {
	noteq := 0
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		noteq++
		nums[left] = nums[i]
		left++
	}

	return noteq
}
