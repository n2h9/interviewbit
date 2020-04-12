package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(1)
	A.left = treeNode_new(2)
	A.left.left = treeNode_new(3)
	A.left.right = treeNode_new(4)
	A.right = treeNode_new(5)
	A.right.left = treeNode_new(6)
	A.right.right = treeNode_new(7)

	expected := []int{3, 2, 4, 1, 6, 5, 7}

	assert.Equal(t, expected, inorderTraversal(A))
}
