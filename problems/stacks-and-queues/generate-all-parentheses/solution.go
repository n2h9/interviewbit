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

func isValid(A string) int {
	if len(A) <= 0 {
		return 1
	}
	var stack *listNode
	for i := 0; i < len(A); i++ {
		b := A[i]
		if isOpen(b) {
			stack = push(stack, b)
			continue
		}
		if isClose(b) {
			if isEmpty(stack) {
				// found close parentheses on empty stack
				return 0
			}
			var v byte
			stack, v = pop(stack)
			if !isCorrectPair(v, b) {
				return 0
			}
			continue
		}
		// not open and not closed byte
		return 0
	}
	if isEmpty(stack) {
		return 1
	}
	return 0

}

func isEmpty(stack *listNode) bool {
	return stack == nil
}

func push(stack *listNode, v byte) *listNode {
	node := listNode_new(int(v))
	node.next = stack
	return node
}

func pop(stack *listNode) (*listNode, byte) {
	if isEmpty(stack) {
		return stack, 0
	}

	v := byte(stack.value)
	node := stack.next
	stack.next = nil
	return node, v
}

func isOpen(c byte) bool {
	return c == '(' || c == '[' || c == '{'
}

func isClose(c byte) bool {
	return c == ')' || c == ']' || c == '}'
}

func isCorrectPair(co, cc byte) bool {
	return co == '(' && cc == ')' ||
		co == '[' && cc == ']' ||
		co == '{' && cc == '}'
}
