package solution

import "fmt"

const DEBUG_ENABLED = false

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

func lca(A *treeNode, B int, C int) int {
	ca, xfound, yfound := find(A, B, C)
	if xfound && yfound {
		return ca
	}
	return -1
}

func find(t *treeNode, x, y int) (int, bool, bool) {
	if t == nil {
		debug("t is nil")
		return 0, false, false
	}
	if isLeaf(t) {
		debug(fmt.Sprintf("t is leave: t.value = %d x = %d y = %d", t.value, x, y))
		if t.value == x && t.value == y {
			return t.value, true, true
		}
		if t.value == x {
			return 0, true, false
		}
		if t.value == y {
			return 0, false, true
		}
		return 0, false, false
	}
	ancestorValue, isXFound, isYFound := find(t.left, x, y)
	debug(fmt.Sprintf("left ancestor value: t.value = %d x = %d y = %d; av = %d xF = %t yF = %t", t.value, x, y, ancestorValue, isXFound, isYFound))
	if isXFound && isYFound {
		return ancestorValue, true, true
	}
	rightAncestorValue, isXFoundRight, isYFoundRight := find(t.right, x, y)
	debug(fmt.Sprintf("right ancestor value: t.value = %d x = %d y = %d; av = %d xF = %t yF = %t", t.value, x, y, rightAncestorValue, isXFoundRight, isYFoundRight))
	if isXFoundRight && isYFoundRight {
		return rightAncestorValue, true, true
	}
	isXFound = isXFound || isXFoundRight || t.value == x
	isYFound = isYFound || isYFoundRight || t.value == y
	if isXFound && isYFound {
		return t.value, true, true
	}
	if isXFound {
		return 0, true, false
	}
	if isYFound {
		return 0, false, true
	}

	return 0, false, false
}

func isLeaf(t *treeNode) bool {
	return t.left == nil && t.right == nil
}

func debug(s string) {
	if DEBUG_ENABLED {
		fmt.Println(s)
	}
}
