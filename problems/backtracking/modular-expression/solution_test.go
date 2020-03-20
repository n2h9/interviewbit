package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 2
	B := 3
	C := 3
	expected := 2
	assert.Equal(t, expected, Mod(A, B, C))
}

func TestSol2(t *testing.T) {
	A := 5
	B := 15
	C := 17
	expected := 7
	assert.Equal(t, expected, Mod(A, B, C))
}

func TestSol3(t *testing.T) {
	A := -1
	B := 1
	C := 20
	expected := 19
	assert.Equal(t, expected, Mod(A, B, C))
}

func TestSol4(t *testing.T) {
	A := 0
	B := 0
	C := 1
	expected := 0
	assert.Equal(t, expected, Mod(A, B, C))
}
