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

// return max_rectangle
func largestRectangleArea(A []int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return A[0]
	}
	// n * 1
	maxSquare := n
	stack := newStack()
	stack = push(stack, 0)
	for i := 1; i < n; i++ {
		if isEmpty(stack) {
			stack = push(stack, i)
			continue
		}
		if A[i] > A[i-1] {
			stack = push(stack, i)
			continue
		}
		if A[i] < A[i-1] {
			var vlast int
			for !isEmpty(stack) {
				v := top(stack)
				if A[v] < A[i] {
					break
				}
				s := (i - v) * A[v]
				if s > maxSquare {
					maxSquare = s
				}
				s = (i - v - 1) * A[vlast]
				if s > maxSquare {
					maxSquare = s
				}
				stack = pop(stack)
				vlast = v
			}
			var s int
			if isEmpty(stack) {
				s = i * A[vlast]
			}
			if s > maxSquare {
				maxSquare = s
			}
			stack = push(stack, i)
			continue
		}
		// the only case left is when A[i] == A[i-1]
		// go further without any actions
	}

	// process all elemets left in stack
	for !isEmpty(stack) {
		v := top(stack)
		s := (n - v) * A[v]
		if s > maxSquare {
			maxSquare = s
		}
		stack = pop(stack)
	}

	return maxSquare
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
		return stack
	}
	node := stack.next
	stack.next = nil
	return node
}

func top(stack *listNode) int {
	return stack.value
}
