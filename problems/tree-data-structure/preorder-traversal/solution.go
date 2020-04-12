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

func preorderTraversal(A *treeNode) []int {
	if A == nil {
		return []int{}
	}

	res := []int{}
	stack := stackNode_new(A)
	for !stack_isEmpty(stack) {
		node := stack_top(stack)
		stack = stack_pop(stack)
		res = append(res, node.value)
		if node.right != nil {
			stack = stack_push(stack, stackNode_new(node.right))
		}
		if node.left != nil {
			stack = stack_push(stack, stackNode_new(node.left))
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
