package main

import "math"

var minDiff, prev int

func getMinimumDifferenceWithPrev(root *TreeNode) int {
	prev = math.MaxInt64
	minDiff = math.MaxInt64
	inorder(root)
	return minDiff
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)

	if root.Val > prev {
		minDiff = min(minDiff, root.Val-prev)
	} else {
		minDiff = min(minDiff, prev-root.Val)
	}
	prev = root.Val

	inorder(root.Right)
	return
}
