package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(1)
	A.left = treeNode_new(2)
	A.right = treeNode_new(3)

	B := treeNode_new(1)
	B.left = treeNode_new(2)
	B.right = treeNode_new(3)
	expected := 1

	assert.Equal(t, expected, isSameTree(A, B))
}

func TestSol2(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(2)
	A.left.left = treeNode_new(1)

	B := treeNode_new(1)
	B.left = treeNode_new(2)
	B.right = treeNode_new(3)
	expected := 0

	assert.Equal(t, expected, isSameTree(A, B))
}
