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
func subtract(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	n := lenList(A)
	isOdd := n&1 != 0
	half := n >> 1
	middleElement := byIndex(A, half-1)
	A2 := middleElement.next
	middleElement.next = nil
	A = revert(A)
	node := A
	node2 := A2
	if isOdd {
		node2 = node2.next
	}
	for node != nil && node2 != nil {
		node.value = node2.value - node.value
		node = node.next
		node2 = node2.next
	}
	endOfA := A
	A = revert(A)
	endOfA.next = A2
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
