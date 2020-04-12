package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(1)
	A.left = treeNode_new(2)

	expected := 2

	assert.Equal(t, expected, maxDepth(A))
}

func TestSol2(t *testing.T) {
	A := treeNode_new(1)
	A.left = treeNode_new(2)
	A.right = treeNode_new(3)
	A.right.right = treeNode_new(4)
	A.right.right.right = treeNode_new(5)

	expected := 4

	assert.Equal(t, expected, maxDepth(A))
}
