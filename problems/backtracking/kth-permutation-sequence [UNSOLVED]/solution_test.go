package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	n := 3
	k := 4
	expected := "231"
	assert.Equal(t, expected, getPermutation(n, k))
}

func TestSol2(t *testing.T) {
	n := 11
	k := 1
	expected := "1234567891011"
	assert.Equal(t, expected, getPermutation(n, k))
}

func TestSol3(t *testing.T) {
	n := 3
	k := 5
	expected := "312"
	assert.Equal(t, expected, getPermutation(n, k))
}
