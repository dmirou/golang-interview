package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type levelSum struct {
	elementsCount int
	sum           float64
}

func averageOfLevelsDFS(root *TreeNode) []float64 {
	sums := make([]levelSum, 1)
	sums = sumOfLevelDFS(root, sums, 0)

	average := make([]float64, len(sums))
	for i, levelSum := range sums {
		average[i] = levelSum.sum / float64(levelSum.elementsCount)
	}

	return average
}

// sumOfLevelDFS cals sum for each level using depth-first search (DFS) обход в глубину
func sumOfLevelDFS(root *TreeNode, sums []levelSum, level int) []levelSum {
	if root == nil {
		return sums
	}

	if len(sums) <= level {
		sums = append(sums, levelSum{})
	}
	sums[level].sum += float64(root.Val)
	sums[level].elementsCount++

	sums = sumOfLevelDFS(root.Left, sums, level+1)

	return sumOfLevelDFS(root.Right, sums, level+1)
}
