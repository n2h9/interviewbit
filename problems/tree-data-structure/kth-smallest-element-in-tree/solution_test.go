package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := treeNode_new(2)
	A.left = treeNode_new(1)
	A.right = treeNode_new(3)
	B := 2

	expected := 2

	assert.Equal(t, expected, kthsmallest(A, B))
}
