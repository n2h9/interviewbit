package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{3, 6, 9, 12}
	expected := 4

	assert.Equal(t, expected, solve(A))
}

func TestSol1(t *testing.T) {
	A := []int{9, 4, 7, 2, 10}
	expected := 3

	assert.Equal(t, expected, solve(A))
}

func TestSol2(t *testing.T) {
	A := []int{3, 3, 3, 3, 3}
	expected := 5

	assert.Equal(t, expected, solve(A))
}

func TestSol3(t *testing.T) {
	A := []int{3, 4, 5, 8, 12, 10, 16, 20, 1, 24, 1, 28, 21}
	expected := 7

	assert.Equal(t, expected, solve(A))
}

func TestSol4(t *testing.T) {
	A := []int{3, 4, 1, 8, 2, 12, 5, 16, 100, 20, 1000, 100000, 24, 28}
	expected := 7

	assert.Equal(t, expected, solve(A))
}
