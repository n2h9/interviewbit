package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(9)
	A.right = treeNode_new(20)
	A.right.left = treeNode_new(15)
	A.right.right = treeNode_new(7)

	expected := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	}

	assert.Equal(t, expected, zigzagLevelOrder(A))
}

func TestSol2(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(9)
	A.left.left = treeNode_new(16)
	A.left.right = treeNode_new(8)
	A.right = treeNode_new(20)
	A.right.left = treeNode_new(15)
	A.right.right = treeNode_new(7)

	expected := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{16, 8, 15, 7},
	}

	assert.Equal(t, expected, zigzagLevelOrder(A))
}

func TestSol3(t *testing.T) {
	A := treeNode_new(3)
	A.left = treeNode_new(9)
	A.left.right = treeNode_new(8)
	A.right = treeNode_new(20)
	A.right.left = treeNode_new(15)

	expected := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{8, 15},
	}

	assert.Equal(t, expected, zigzagLevelOrder(A))
}
