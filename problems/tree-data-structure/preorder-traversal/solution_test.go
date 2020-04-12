package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(1)
	A.right = treeNode_new(2)
	A.right.left = treeNode_new(3)

	expected := []int{1, 2, 3}

	assert.Equal(t, expected, preorderTraversal(A))
}
