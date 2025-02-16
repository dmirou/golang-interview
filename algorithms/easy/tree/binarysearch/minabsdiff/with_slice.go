package main

func getMinimumDifferenceWithSlice(root *TreeNode) int {
	sl := toSlice(root)
	diff := sl[1] - sl[0]
	for i := 2; i < len(sl); i++ {
		diff = min(diff, sl[i]-sl[i-1])
	}

	return diff
}

func toSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, toSlice(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, toSlice(root.Right)...)
	}

	return res
}
