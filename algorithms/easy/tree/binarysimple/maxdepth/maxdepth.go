package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	st := newStack()

	push := func(node *TreeNode, depth int) {
		st.push(&depthNode{node: node, depth: depth})
	}

	maxDepth := 0
	push(root, 1)

	for !st.empty() {
		cur := st.pop()
		if maxDepth < cur.depth {
			maxDepth = cur.depth
		}

		if cur.node.Left != nil {
			push(cur.node.Left, cur.depth+1)
		}
		if cur.node.Right != nil {
			push(cur.node.Right, cur.depth+1)
		}
	}

	return maxDepth
}

type depthNode struct {
	node  *TreeNode
	depth int
}

type stack struct {
	elements []*depthNode
}

func newStack() stack {
	return stack{elements: make([]*depthNode, 0)}
}

func (s *stack) push(el *depthNode) {
	s.elements = append(s.elements, el)
}

func (s *stack) pop() *depthNode {
	if len(s.elements) == 0 {
		return nil
	}

	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return last
}

func (s *stack) empty() bool {
	return len(s.elements) == 0
}
