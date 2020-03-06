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
func insertionSortList(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}

	head := A
	prev := head
	node := head.next
	for node != nil {
		next := node.next
		nodeSorted := head
		var prevSorted *listNode = nil
		for nodeSorted != nil && nodeSorted != node {
			if node.value >= nodeSorted.value {
				// go futher
				prevSorted = nodeSorted
				nodeSorted = nodeSorted.next
				continue
			}
			// insert
			prev.next = next
			node.next = nodeSorted
			if prevSorted == nil {
				head = node
			} else {
				prevSorted.next = node
			}
			break
			// --
		}
		if node.next == next {
			// node was not moved
			prev = node
			node = next
		} else {
			// prev is already set when node was moved on its place
			node = next
		}
	}
	return head
}
