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
	Ain := []int{1, 2, 3, 4, 5}
	Aout := []int{5, 4, 3, 2, 1}
	res := listToSlice(
		reverseList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol2(t *testing.T) {
	Ain := []int{1, 2, 3, 4}
	Aout := []int{4, 3, 2, 1}
	res := listToSlice(
		reverseList(
			sliceToList(Ain),
		),
	)
	assert.Equal(t, Aout, res)
}
