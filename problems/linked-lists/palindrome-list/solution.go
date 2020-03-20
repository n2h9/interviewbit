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
 * @Output Integer
 */
func lPalin(A *listNode) int {
	if A == nil || A.next == nil {
		return 1
	}
	n := lenList(A)
	half2Index := n >> 1
	// equal to n % 2 != 0
	if n&1 == 1 {
		half2Index++
	}
	halfList := byIndex(A, half2Index)
	halfList = revert(halfList)
	// important to iterate over second half,
	// because first half still has pointer to element from second half
	for halfList != nil && A != nil {
		if halfList.value != A.value {
			return 0
		}
		halfList = halfList.next
		A = A.next
	}

	return 1
}

func lenList(l *listNode) int {
	c := 0
	for ; l != nil; l = l.next {
		c++
	}
	return c
}

func byIndex(l *listNode, i int) *listNode {
	for ; i > 0; i-- {
		l = l.next
	}
	return l
}

func revert(l *listNode) *listNode {
	if l == nil || l.next == nil {
		return l
	}
	prev := l
	l = l.next
	prev.next = nil
	for l != nil {
		next := l.next
		l.next = prev
		prev = l
		l = next
	}
	return prev
}
