package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	/*
					10
			5				20
		1		8		19		25
	*/

	root := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 19,
			},
			Right: &TreeNode{
				Val: 25,
			},
		},
	}

	fmt.Println(getMinimumDifferenceWithSlice(&root))
	fmt.Println(getMinimumDifferenceWithPrev(&root))

	// [236,104,701,null,227,null,911]
	//						236
	//			104					701
	//				227						911
	root = TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}
	fmt.Println(getMinimumDifferenceWithSlice(&root))
	fmt.Println(getMinimumDifferenceWithPrev(&root))

	// [0,null,100000]
	// 			0
	//					100000
	root = TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: 100000,
		},
	}
	fmt.Println(getMinimumDifferenceWithSlice(&root))
	fmt.Println(getMinimumDifferenceWithPrev(&root))
}
