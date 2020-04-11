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

func isSameTree(A *treeNode, B *treeNode) int {
	if A == nil && B == nil {
		return 1
	}
	if A != nil && B == nil {
		return 0
	}
	if B != nil && A == nil {
		return 0
	}
	if A.value != B.value {
		return 0
	}
	if v := isSameTree(A.left, B.left); v == 0 {
		return 0
	}

	return isSameTree(A.right, B.right)
}
