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

func invertTree(A *treeNode) *treeNode {
	if A == nil {
		return A
	}
	left := A.left
	A.left = invertTree(A.right)
	A.right = invertTree(left)
	return A
}
