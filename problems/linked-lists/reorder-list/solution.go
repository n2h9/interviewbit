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
func reorderList(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	lh := lenList(A) / 2
	m := byIndex(A, lh-1)
	A2 := m.next
	// to nrake the loop futher
	m.next = nil
	A2 = revert(A2)

	node := A
	node2 := A2
	last := node2
	for node != nil {
		next := node.next
		next2 := node2.next
		node.next = node2
		node2.next = next

		node = next
		last = node2
		node2 = next2
	}
	if node2 != nil {
		last.next = node2
	}

	return A
}

func lenList(l *listNode) int {
	if l == nil {
		return 0
	}
	c := 1
	for ; l.next != nil; l = l.next {
		c++
	}
	return c
}

func byIndex(l *listNode, i int) *listNode {

	m := l
	for ; i > 0; i-- {
		m = m.next
	}
	return m
}

func revert(l *listNode) *listNode {
	if l == nil {
		return nil
	}
	curr := l.next
	l.next = nil
	for curr != nil {
		next := curr.next
		curr.next = l
		l = curr
		curr = next
	}

	return l
}
