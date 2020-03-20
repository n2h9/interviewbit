package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{4, 5, 2, 10, 8}
	expected := []int{-1, 4, -1, 2, 2}

	assert.Equal(t, expected, prevSmaller(A))
}

func TestSol2(t *testing.T) {
	A := []int{3, 2, 1}
	expected := []int{-1, -1, -1}

	assert.Equal(t, expected, prevSmaller(A))
}

func TestSol3(t *testing.T) {
	A := []int{1, 2, 3}
	expected := []int{-1, 1, 2}

	assert.Equal(t, expected, prevSmaller(A))
}
