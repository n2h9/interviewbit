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

func zigzagLevelOrder(A *treeNode) [][]int {
	if A == nil {
		return [][]int{}
	}

	stack := stackNode_new(A)
	stack1 := stackNode_newEmpty()
	order := ORDER_LR
	res := make([][]int, 1)
	i := 0

	for !stack_isEmpty(stack) {
		t := stack_top(stack)
		stack = stack_pop(stack)
		if i == len(res) {
			res = append(res, nil)
		}
		res[i] = append(res[i], t.value)
		switch {
		case isLeaf(t):
			// do nothing :)
		case t.left == nil:
			stack1 = stack_push(stack1, stackNode_new(t.right))
		case t.right == nil:
			stack1 = stack_push(stack1, stackNode_new(t.left))
		case order == ORDER_LR:
			stack1 = stack_push(stack1, stackNode_new(t.left))
			stack1 = stack_push(stack1, stackNode_new(t.right))
		case order == ORDER_RL:
			stack1 = stack_push(stack1, stackNode_new(t.right))
			stack1 = stack_push(stack1, stackNode_new(t.left))
		}
		if stack_isEmpty(stack) {
			stack = stack1
			stack1 = nil
			order = changeOrder(order)
			i++
		}
	}

	return res
}

func isLeaf(t *treeNode) bool {
	return t.left == nil && t.right == nil
}

type order bool

const ORDER_LR order = true
const ORDER_RL order = false

func changeOrder(val order) order {
	if val == ORDER_LR {
		return ORDER_RL
	}
	return ORDER_LR
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

func stackNode_newEmpty() *stackNode {
	return nil
}

func stack_push(stack *stackNode, node *stackNode) *stackNode {
	if stack_isEmpty(stack) {
		return node
	}
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
