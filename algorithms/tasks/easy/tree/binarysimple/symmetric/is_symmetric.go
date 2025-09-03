package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSym(root.Left, root.Right)
}

func isSym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
}
