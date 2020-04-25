package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{1, 2, 1, 3, 4, 3}
	B := 3
	expected := []int{2, 3, 3, 2}

	assert.Equal(t, expected, dNums(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{7, 6, 4, 2, 1}
	B := 2
	expected := []int{2, 2, 2, 2}

	assert.Equal(t, expected, dNums(A, B))
}

func TestSol3(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	B := 4
	expected := []int{1, 1}

	assert.Equal(t, expected, dNums(A, B))
}

func TestSol4(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	B := 10
	expected := []int{}

	assert.Equal(t, expected, dNums(A, B))
}
