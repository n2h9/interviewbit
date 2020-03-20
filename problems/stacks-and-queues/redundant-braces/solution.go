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

func braces(A string) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	var stack *listNode
	for i := 0; i < n; i++ {
		if !isCloseBrace(A[i]) {
			// add all stuff to stack
			stack = push(stack, int(A[i]))
			continue
		}

		// if close brace => restore all values until open brace
		// and determine if inside brace
		// was at least one operation, this is a marker of non redundant brace
		bracesAreRedundant := true
		for !isEmpty(stack) {
			val := byte(top(stack))
			stack = pop(stack)
			if isOpenBrace(val) {
				break
			}
			if isOperation(val) {
				bracesAreRedundant = false
			}
		}
		// if at least one braces are redundant stop and return
		if bracesAreRedundant {
			return 1
		}
	}

	return 0
}

func isOperation(ch byte) bool {
	for _, v := range []byte{'+', '-', '*', '/'} {
		if ch == v {
			return true
		}
	}

	return false
}

func isOpenBrace(ch byte) bool {
	return ch == '('
}

func isCloseBrace(ch byte) bool {
	return ch == ')'
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
