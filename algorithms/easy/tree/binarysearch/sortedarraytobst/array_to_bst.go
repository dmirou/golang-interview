package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1 8 10] => [8 1 10]
// [1 8] => [8 1 null]
// [1 2 3 4] => [3 2 4 1 null null null]
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var root TreeNode

	middle := len(nums) / 2
	root.Val = nums[middle]
	if len(nums) == 1 {
		return &root
	}

	root.Left = sortedArrayToBST(nums[:middle])
	if middle+1 <= len(nums)-1 {
		root.Right = sortedArrayToBST(nums[middle+1:])
	}

	return &root
}

func main() {

}
