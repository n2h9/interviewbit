package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := listNode_new(10)
	assert.Equal(t, A, insertionSortList(A))
}
