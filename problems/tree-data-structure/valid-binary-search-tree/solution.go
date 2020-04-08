package solution

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
 * @Output Integer
 */
// func isValidBST(A *treeNode) int {
// 	if A == nil {
// 		return 1
// 	}

// 	// check left child correctnes
// 	if A.left != nil && A.left.value > A.value {
// 		return 0
// 	}

// 	// check right child correctnes
// 	if A.right != nil && A.right.value <= A.value {
// 		return 0
// 	}

// 	// check left sutreee correctnes
// 	if v := isValidBST(A.left); v == 0 {
// 		return 0
// 	}

// 		// check right sutreee correctnes
// 		return isValidBST(A.right)
// }

func isValidBST(A *treeNode) int {
	return validateWithBorders(A, nil, nil)
}

func validateWithBorders(A *treeNode, bLeft, bRight *int) int {
	if A == nil {
		return 1
	}
	// 1) validate value of current node:
	// value is out of left border
	if bLeft != nil && A.value <= *bLeft {
		return 0
	}
	// value is out of right border
	if bRight != nil && A.value >= *bRight {
		return 0
	}

	// 2) validate left subtree
	if v := validateWithBorders(A.left, bLeft, &A.value); v == 0 {
		return 0
	}

	return validateWithBorders(A.right, &A.value, bRight)
}
