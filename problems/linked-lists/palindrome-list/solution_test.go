package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sliceToList(A []int) *listNode {
	var head, prev *listNode
	for _, v := range A {
		val := listNode_new(v)
		if head == nil {
			head = val
			prev = val
			continue
		}
		prev.next = val
		prev = val
	}
	return head
}

func TestSol(t *testing.T) {
	Ain := []int{1, 2, 1}
	res := lPalin(
		sliceToList(Ain),
	)
	assert.Equal(t, 1, res)
}

func TestSol2(t *testing.T) {
	Ain := []int{1, 2, 3}
	res := lPalin(
		sliceToList(Ain),
	)
	assert.Equal(t, 0, res)
}

func TestSol3(t *testing.T) {
	Ain := []int{1}
	res := lPalin(
		sliceToList(Ain),
	)
	assert.Equal(t, 1, res)
}

func TestSol4(t *testing.T) {
	Ain := []int{1, 2, 3, 4, 4, 3, 2, 1}
	res := lPalin(
		sliceToList(Ain),
	)
	assert.Equal(t, 1, res)
}

func TestSol5(t *testing.T) {
	Ain := []int{1, 2, 3, 4, 5, 3, 2, 1}
	res := lPalin(
		sliceToList(Ain),
	)
	assert.Equal(t, 0, res)
}
