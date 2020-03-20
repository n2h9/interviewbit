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
 * @input B : Integer
 * @input C : Integer
 *
 * @Output head pointer of list.
 */
func reverseBetween(A *listNode, B int, C int) *listNode {
	if A == nil || A.next == nil || C <= B {
		return A
	}
	C -= B

	// the element before element #B
	var sentryLeft *listNode
	B-- // first iteration is above
	if B > 0 {
		// to avoid compare sentryLeft with nil in cycle
		sentryLeft = A
		B--
	}
	for ; B > 0; B-- {
		// swipe to position B
		sentryLeft = sentryLeft.next
	}
	node := A
	if sentryLeft != nil {
		node = sentryLeft.next
	}
	sentryRight := node
	prev := node
	node = node.next

	for ; C > 0; C-- {
		next := node.next
		node.next = prev
		prev = node
		node = next
	}
	sentryRight.next = node
	if sentryLeft == nil {
		A = prev
	} else {
		sentryLeft.next = prev
	}
	return A
}
