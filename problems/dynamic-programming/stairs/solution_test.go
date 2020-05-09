package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 2
	expected := 2

	assert.Equal(t, expected, climbStairs(A))
}

func TestSol1(t *testing.T) {
	A := 3
	expected := 3

	assert.Equal(t, expected, climbStairs(A))
}

func TestSol2(t *testing.T) {
	A := 4
	expected := 5

	assert.Equal(t, expected, climbStairs(A))
}

func TestSol3(t *testing.T) {
	A := 5
	expected := 8

	assert.Equal(t, expected, climbStairs(A))
}
