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

func inorderTraversal(A *treeNode) []int {
	if A == nil {
		return []int{}
	}

	res := []int{}
	processedMap := make(map[*treeNode]bool)
	stack := stackNode_new(A)
	for !stack_isEmpty(stack) {
		node := stack_top(stack)
		if node.left != nil && !processedMap[node.left] {
			stack = stack_push(stack, stackNode_new(node.left))
			continue
		}

		res = append(res, node.value)
		processedMap[node] = true
		stack = stack_pop(stack)
		if node.right != nil {
			stack = stack_push(stack, stackNode_new(node.right))
		}
	}
	return res
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
