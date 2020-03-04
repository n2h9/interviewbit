package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{2, 1, 4, 3, 2}
	B := 3
	expected := 2

	assert.Equal(t, expected, kthsmallest(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{3, 3, 3, 3, 3}
	B := 3
	expected := 3

	assert.Equal(t, expected, kthsmallest(A, B))
}

func TestSol3(t *testing.T) {
	A := []int{20000, 20000, 200000, 200000, 200000}
	B := 3
	expected := 200000

	assert.Equal(t, expected, kthsmallest(A, B))
}

func TestSol4(t *testing.T) {
	A := []int{60, 94, 63, 3, 86, 40, 93, 36, 56, 48, 17, 10, 23, 43, 77, 1, 1, 93, 79, 4, 10, 47, 1, 99, 91, 53, 99, 18, 52, 61, 84, 10, 13, 52, 3, 9, 78, 16, 7, 62}
	B := 22
	expected := 52

	assert.Equal(t, expected, kthsmallest(A, B))
}
