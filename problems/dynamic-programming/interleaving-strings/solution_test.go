package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := "aabcc"
	B := "dbbca"
	C := "aadbbcbcac"
	expected := 1

	assert.Equal(t, expected, isInterleave(A, B, C))
}

func TestSol1(t *testing.T) {
	A := "aabcc"
	B := "dbbca"
	C := "aadbbbaccc"
	expected := 0

	assert.Equal(t, expected, isInterleave(A, B, C))
}

func TestSol2(t *testing.T) {
	A := "aabccffuuuuu"
	B := "dbbcazz"
	C := "aadbbcbcaczfzfuuuuu"
	expected := 1

	assert.Equal(t, expected, isInterleave(A, B, C))
}

func TestSol3(t *testing.T) {
	A := "aabccffuuuuu"
	B := "dbbcazzq"
	C := "aadbbcbcaczfzfuuuuu"
	expected := 0

	assert.Equal(t, expected, isInterleave(A, B, C))
}

func TestSol4(t *testing.T) {
	A := "q2idqciLUL6mVGR"
	B := "9PdcVkqtgHJ3p"
	C := "q9Pd2cVikdqcqitLUgLHJ63pmVGR"
	expected := 1

	assert.Equal(t, expected, isInterleave(A, B, C))
}
