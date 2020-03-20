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
	A := []int{2, 4, 3}
	B := []int{5, 6, 4}
	expected := []int{7, 0, 8}
	res := listToSlice(
		addTwoNumbers(
			sliceToList(A),
			sliceToList(B),
		),
	)
	assert.Equal(t, expected, res)
}

func TestSol2(t *testing.T) {
	A := []int{9}
	B := []int{5, 6, 4}
	expected := []int{4, 7, 4}
	res := listToSlice(
		addTwoNumbers(
			sliceToList(A),
			sliceToList(B),
		),
	)
	assert.Equal(t, expected, res)
}

func TestSol3(t *testing.T) {
	A := []int{9}
	B := []int{9, 9, 9}
	expected := []int{8, 0, 0, 1}
	res := listToSlice(
		addTwoNumbers(
			sliceToList(A),
			sliceToList(B),
		),
	)
	assert.Equal(t, expected, res)
}

func TestSol4(t *testing.T) {
	A := []int{9, 9, 9}
	B := []int{9, 9, 9}
	expected := []int{8, 9, 9, 1}
	res := listToSlice(
		addTwoNumbers(
			sliceToList(A),
			sliceToList(B),
		),
	)
	assert.Equal(t, expected, res)
}
