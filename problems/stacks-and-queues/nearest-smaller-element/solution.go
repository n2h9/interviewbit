package solution

type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

func prevSmaller(A []int) []int {
	n := len(A)
	if n <= 0 {
		return A
	}
	stack := newStack()
	B := make([]int, n)
	for i := 0; i < n; i++ {
		if isEmpty(stack) {
			B[i] = -1
			stack = push(stack, A[i])
			continue
		}
		var v int
		for !isEmpty(stack) {
			v = top(stack)
			if v < A[i] {
				break
			}
			stack = pop(stack)
		}
		if isEmpty(stack) {
			v = -1
		}
		B[i] = v
		stack = push(stack, A[i])
	}
	return B
}

func newStack() *listNode {
	return nil
}

func isEmpty(stack *listNode) bool {
	return stack == nil
}

func push(stack *listNode, v int) *listNode {
	node := listNode_new(v)
	node.next = stack
	return node
}

func pop(stack *listNode) *listNode {
	if isEmpty(stack) {
		return nil
	}
	node := stack.next
	stack.next = nil
	return node
}

func top(stack *listNode) int {
	if isEmpty(stack) {
		return -1
	}
	return stack.value
}
