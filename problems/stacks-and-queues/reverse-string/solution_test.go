package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := "abc"
	expected := "cba"

	assert.Equal(t, expected, reverseString(A))
}

func TestSol2(t *testing.T) {
	A := ""
	expected := ""

	assert.Equal(t, expected, reverseString(A))
}

func TestSol3(t *testing.T) {
	A := "Кирилл"
	expected := "ллириК"

	assert.Equal(t, expected, reverseString(A))
}
