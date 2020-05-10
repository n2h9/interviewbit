package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{1, 2}
	B := []int{2, 1}
	expected := 1

	assert.Equal(t, expected, canCompleteCircuit(A, B))
}

func TestSol1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{3, 4, 5, 1, 2}
	expected := 3

	assert.Equal(t, expected, canCompleteCircuit(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{2, 3, 4}
	B := []int{3, 4, 3}
	expected := -1

	assert.Equal(t, expected, canCompleteCircuit(A, B))
}
