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

func kthsmallest(A *treeNode, B int) int {
	if A == nil {
		return -1
	}

	stack := stackNode_new(A)
	visitedNodesMap := make(map[*treeNode]bool)
	val := -1

	for !stack_isEmpty(stack) && B > 0 {
		t := stack_top(stack)
		//1) put left node in a stack
		if t.left != nil && !visitedNodesMap[t.left] {
			stack = stack_push(stack, stackNode_new(t.left))
			continue
		}
		//2) process node
		stack = stack_pop(stack)
		visitedNodesMap[t] = true
		val = t.value
		B--

		// 3) push right node
		if t.right != nil {
			stack = stack_push(stack, stackNode_new(t.right))
		}
	}

	return val
}

type stackNode struct {
	value *treeNode
	next  *stackNode
}

func stackNode_new(val *treeNode) *stackNode {
	return &stackNode{
		value: val,
	}
}

func stack_push(stack *stackNode, node *stackNode) *stackNode {
	node.next = stack
	return node
}

func stack_pop(stack *stackNode) *stackNode {
	next := stack.next
	stack.next = nil
	return next
}

func stack_top(stack *stackNode) *treeNode {
	return stack.value
}

func stack_isEmpty(stack *stackNode) bool {
	return stack == nil
}
