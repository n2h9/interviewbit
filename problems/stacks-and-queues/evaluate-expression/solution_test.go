package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []string{"2", "1", "+", "3", "*"}
	expected := 9
	assert.Equal(t, expected, evalRPN(A))
}

func TestSol2(t *testing.T) {
	A := []string{"4", "13", "5", "/", "+"}
	expected := 6
	assert.Equal(t, expected, evalRPN(A))
}

func TestSol3(t *testing.T) {
	A := []string{"15", "8", "5", "-", "*", "-3", "/"}
	expected := -15
	assert.Equal(t, expected, evalRPN(A))
}
