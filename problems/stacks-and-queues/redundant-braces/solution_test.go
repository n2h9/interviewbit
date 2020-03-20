package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := "((a + b))"
	expected := 1

	assert.Equal(t, expected, braces(A))
}

func TestSol2(t *testing.T) {
	A := "(a + (a + b))"
	expected := 0

	assert.Equal(t, expected, braces(A))
}

func TestSol3(t *testing.T) {
	A := "(a + (a + b) + (z + x))"
	expected := 0

	assert.Equal(t, expected, braces(A))
}

func TestSol4(t *testing.T) {
	A := "((a+b)+(a+b))"
	expected := 0

	assert.Equal(t, expected, braces(A))
}
