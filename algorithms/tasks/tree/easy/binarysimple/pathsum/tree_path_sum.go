package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(root, 0, targetSum)
}

func pathSum(node *TreeNode, curSum, targetSum int) bool {
	if node == nil {
		return false
	}

	curSum += node.Val
	if curSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	return pathSum(node.Left, curSum, targetSum) ||
		pathSum(node.Right, curSum, targetSum)
}
