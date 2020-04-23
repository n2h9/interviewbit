package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 3
	B := []int{6, 5}
	expected := 14

	assert.Equal(t, expected, nchoc(A, B))
}

func TestSol2(t *testing.T) {
	A := 10
	B := []int{2147483647, 2000000014, 2147483647}
	expected := 284628164

	assert.Equal(t, expected, nchoc(A, B))
}
