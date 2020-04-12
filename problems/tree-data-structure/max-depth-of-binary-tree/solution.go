package solution

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

func maxDepth(A *treeNode) int {
	if A == nil {
		return 0
	}
	if isLeaf(A) {
		return 1
	}
	return 1 + max(maxDepth(A.left), maxDepth(A.right))
}

func isLeaf(t *treeNode) bool {
	return t.left == nil && t.right == nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
