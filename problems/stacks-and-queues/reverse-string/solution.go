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

func reverseString(A string) string {
	if len(A) <= 0 {
		return A
	}

	var stack *listNode

	for _, v := range A {
		node := listNode_new(int(v))
		node.next = stack
		stack = node
	}

	res := []rune{}

	for node := stack; node != nil; node = node.next {
		res = append(res, rune(node.value))
	}

	return string(res)
}
