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

func isBalanced(A *treeNode) int {
	debug("=====")
	ok, h := isBalancedWithHeight(A)
	debug(fmt.Sprintf("%t %d", ok, h))
	if !ok {
		return 0
	}
	return 1
}

func isBalancedWithHeight(t *treeNode) (bool, int) {
	if t == nil {
		debug("Node nil")
		return true, -1
	}
	if isLeave(t) {
		debug(fmt.Sprintf("Node %d is leave\n", t.value))
		return true, 0
	}
	leftBalanced, leftHeight := isBalancedWithHeight(t.left)
	debug(fmt.Sprintf("Node %d: left height = %d, is balanced = %t", t.value, leftHeight, leftBalanced))
	if !leftBalanced {
		// do not return height if not balanced
		return false, -100
	}
	rightBalanced, rightHeight := isBalancedWithHeight(t.right)
	debug(fmt.Sprintf("Node %d: right height = %d, is balanced = %t", t.value, rightHeight, rightBalanced))
	if !rightBalanced {
		// do not return height if not balanced
		return false, -100
	}
	height, isLeftMax := max(leftHeight, rightHeight)
	height++
	var ok bool
	// determine |leftHeight-rightHeight| <= 1
	if isLeftMax {
		ok = (leftHeight - rightHeight) <= 1
	} else {
		ok = (rightHeight - leftHeight) <= 1
	}

	return ok, height
}

func isLeave(t *treeNode) bool {
	return t.left == nil && t.right == nil
}

// @return
// (max value, true if a > b)
func max(a, b int) (int, bool) {
	if a > b {
		return a, true
	}
	return b, false
}

func debug(s string) {
	if DEBUG_ENABLED {
		fmt.Println(s)
	}
}
