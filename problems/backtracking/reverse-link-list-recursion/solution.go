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

/**
 * @input A : Head pointer of linked list
 *
 * @Output head pointer of list.
 */
func reverseList(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	first, _ := reverse(A)
	return first
}

func reverse(start *listNode) (first *listNode, last *listNode) {
	if start.next == nil {
		return start, start
	}
	next := start.next
	start.next = nil
	first, last = reverse(next)
	last.next = start
	return first, start
}
