package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(5)
	A.left.left = treeNode_new(6)
	A.left.right = treeNode_new(2)
	A.left.right.left = treeNode_new(7)
	A.left.right.right = treeNode_new(4)
	A.right = treeNode_new(1)
	A.right.left = treeNode_new(0)
	A.right.right = treeNode_new(8)
	B := 5
	C := 1

	expected := 3

	assert.Equal(t, expected, lca(A, B, C))
}

func TestSol2(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(5)
	A.left.left = treeNode_new(6)
	A.left.right = treeNode_new(2)
	A.left.right.left = treeNode_new(7)
	A.left.right.right = treeNode_new(4)
	A.right = treeNode_new(1)
	A.right.left = treeNode_new(0)
	A.right.right = treeNode_new(8)
	B := 5
	C := 4

	expected := 5

	assert.Equal(t, expected, lca(A, B, C))
}

func TestSol3(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(5)
	A.left.left = treeNode_new(6)
	A.left.right = treeNode_new(2)
	A.left.right.left = treeNode_new(7)
	A.left.right.right = treeNode_new(4)
	A.right = treeNode_new(1)
	A.right.left = treeNode_new(0)
	A.right.right = treeNode_new(8)
	B := 5
	C := 17

	expected := -1

	assert.Equal(t, expected, lca(A, B, C))
}

func TestSol4(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(1)
	B := 1
	C := 1

	expected := 1

	assert.Equal(t, expected, lca(A, B, C))
}
