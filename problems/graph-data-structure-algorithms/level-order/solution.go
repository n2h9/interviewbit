package solution

// Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output 2D int array.
 */
func levelOrder(A *treeNode) [][]int {
	if A == nil {
		return [][]int{}
	}
	res := [][]int{}
	q := &queue{}
	queue_enqueue(q, queueNode_new(A, 0))
	for !queue_isEmpty(q) {
		t := queue_top(q)
		q = queue_dequeue(q)
		if t.value.left != nil {
			q = queue_enqueue(q, queueNode_new(t.value.left, t.level+1))
		}
		if t.value.right != nil {
			q = queue_enqueue(q, queueNode_new(t.value.right, t.level+1))
		}
		if t.level >= len(res) {
			res = append(res, []int{})
		}
		res[t.level] = append(res[t.level], t.value.value)
	}
	return res
}

type queueNode struct {
	value *treeNode
	level int
	next  *queueNode
	prev  *queueNode
}

func queueNode_new(t *treeNode, level int) *queueNode {
	return &queueNode{t, level, nil, nil}
}

type queue struct {
	head *queueNode
	tail *queueNode
}

func queue_isEmpty(q *queue) bool {
	return q.head == nil
}

func queue_top(q *queue) *queueNode {
	return q.head
}

func queue_enqueue(q *queue, val *queueNode) *queue {
	if q.head == nil {
		q.head = val
		q.tail = val
		return q
	}

	val.next = q.tail
	q.tail.prev = val
	q.tail = val
	return q
}

func queue_dequeue(q *queue) *queue {
	if q.head == nil {
		return q
	}
	if q.head.prev != nil {
		q.head.prev.next = nil
	}
	q.head = q.head.prev
	if q.head == nil {
		q.tail = nil
	}
	return q
}
