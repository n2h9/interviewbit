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

func addTwoNumbers(A *listNode, B *listNode) *listNode {
	if B == nil {
		return A
	}
	if A == nil {
		return B
	}

	var sum, curr *listNode

	overflow := 0
	for A != nil || B != nil {
		aVal, bVal := 0, 0
		if A != nil {
			aVal = A.value
		}
		if B != nil {
			bVal = B.value
		}
		s := aVal + bVal + overflow
		overflow = 0
		if s >= 10 {
			s -= 10
			overflow = 1
		}
		node := listNode_new(s)
		if curr == nil {
			sum = node
			curr = node
		} else {
			curr.next = node
			curr = curr.next
		}

		if A != nil {
			A = A.next
		}
		if B != nil {
			B = B.next
		}
	}

	if overflow > 0 {
		curr.next = listNode_new(overflow)
	}

	return sum
}
