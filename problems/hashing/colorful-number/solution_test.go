package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 23
	expected := 1

	assert.Equal(t, expected, colorful(A))
}

func TestSol2(t *testing.T) {
	A := 3245
	expected := 1

	assert.Equal(t, expected, colorful(A))
}

func TestSol3(t *testing.T) {
	A := -9
	expected := 1

	assert.Equal(t, expected, colorful(A))
}
