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
		[]int{9, 20},
		[]int{15, 7},
	}

	assert.Equal(t, expected, levelOrder(A))
}
