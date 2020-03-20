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
	B := 2
	C := 4
	Aout := []int{1, 4, 3, 2, 5}
	res := listToSlice(
		reverseBetween(
			sliceToList(Ain),
			B,
			C,
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol2(t *testing.T) {
	Ain := []int{0, 1, 2, 3, 4, 5}
	B := 3
	C := 4
	Aout := []int{0, 1, 3, 2, 4, 5}
	res := listToSlice(
		reverseBetween(
			sliceToList(Ain),
			B,
			C,
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol3(t *testing.T) {
	Ain := []int{0, 1, 2, 3, 4, 5}
	B := 3
	C := 3
	Aout := []int{0, 1, 2, 3, 4, 5}
	res := listToSlice(
		reverseBetween(
			sliceToList(Ain),
			B,
			C,
		),
	)
	assert.Equal(t, Aout, res)
}

func TestSol4(t *testing.T) {
	Ain := []int{1, 2, 3}
	B := 1
	C := 2
	Aout := []int{2, 1, 3}
	res := listToSlice(
		reverseBetween(
			sliceToList(Ain),
			B,
			C,
		),
	)
	assert.Equal(t, Aout, res)
}
