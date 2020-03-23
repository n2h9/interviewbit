package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 1
	expected := []string{"()"}
	assert.Equal(t, expected, generateParenthesis(A))
}

func TestSol2(t *testing.T) {
	A := 2
	expected := []string{"(())", "()()"}
	assert.Equal(t, expected, generateParenthesis(A))
}

func TestSol3(t *testing.T) {
	A := 3
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	assert.Equal(t, expected, generateParenthesis(A))
}

func TestSol4(t *testing.T) {
	A := 4
	expected := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	assert.Equal(t, expected, generateParenthesis(A))
}
