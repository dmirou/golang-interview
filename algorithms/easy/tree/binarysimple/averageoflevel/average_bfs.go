package main

/*
1
2 3
4 5 6 7

queue 1
count 1, sum 0
i = 0; i < 1
sum 1,
queue 1, 2, 3
i = 1, 1 < 1
sums = [1]
queue 2, 3

queue 2, 3
count 2, sum 0
i 0, 0 < 2
sum 2
queue 2, 3, 4, 5
i 1, 1 <2
sum 5
queue 2, 3, 4, 5, 6, 7
i 2, 2 < 2

sums = [1, 2.5]
queue = 4, 5, 6, 7

*/

// averageOfLevelsVBFS cals sum for each level using breadth-first search (BFS) обход в ширину
func averageOfLevelsVBFS(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	sums := make([]float64, 0)

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		count := len(queue)
		sum := 0

		for i := 0; i < count; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		sums = append(sums, float64(sum)/float64(count))
		queue = queue[count:]
	}

	return sums
}
