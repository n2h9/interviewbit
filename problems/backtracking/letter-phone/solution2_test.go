package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2Sol(t *testing.T) {
	A := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, expected, letterCombinationsNoRecursion(A))
}

func Test2Sol2(t *testing.T) {
	A := "123"
	expected := []string{"1ad", "1ae", "1af", "1bd", "1be", "1bf", "1cd", "1ce", "1cf"}
	assert.Equal(t, expected, letterCombinationsNoRecursion(A))
}
