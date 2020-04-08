package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(1)
	A.left = treeNode_new(2)
	A.right = treeNode_new(3)
	expected := 0

	assert.Equal(t, expected, isValidBST(A))
}

func TestSol2(t *testing.T) {
	A := treeNode_new(2)
	A.left = treeNode_new(1)
	A.right = treeNode_new(3)
	expected := 1

	assert.Equal(t, expected, isValidBST(A))
}

func TestSol3(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(1)
	A.right = treeNode_new(2)
	expected := 0

	assert.Equal(t, expected, isValidBST(A))
}

func TestSol4(t *testing.T) {
	A := treeNode_new(11)
	A.left = treeNode_new(4)
	A.right = treeNode_new(2)
	A.left.left = treeNode_new(5)
	A.left.right = treeNode_new(1)
	A.right.left = treeNode_new(5)
	expected := 0

	assert.Equal(t, expected, isValidBST(A))
}

func TestSol5(t *testing.T) {
	A := treeNode_new(4)
	A.left = treeNode_new(2)
	A.right = treeNode_new(5)
	A.left.left = treeNode_new(1)
	A.left.right = treeNode_new(5)
	expected := 0

	assert.Equal(t, expected, isValidBST(A))
}
