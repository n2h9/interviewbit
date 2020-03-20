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

func listToSlice(l *listNode) []int {
	A := []int{}
	for k := l; k != nil; k = k.next {
		A = append(A, k.value)
	}
	return A
}

func TestSol(t *testing.T) {
	Ain := []int{1, 2, 3, 4}
	Aout := []int{1, 4, 2, 3}
	res := listToSlice(
		reorderList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol2(t *testing.T) {
	Ain := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Aout := []int{1, 8, 2, 7, 3, 6, 4, 5}
	res := listToSlice(
		reorderList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol3(t *testing.T) {
	Ain := []int{1, 2, 3, 4, 5, 6, 7}
	Aout := []int{1, 7, 2, 6, 3, 5, 4}
	res := listToSlice(
		reorderList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol4(t *testing.T) {
	Ain := []int{23, 12, 96}
	Aout := []int{23, 96, 12}
	res := listToSlice(
		reorderList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}
