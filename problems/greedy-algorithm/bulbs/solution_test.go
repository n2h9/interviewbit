package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{0, 1, 0, 1}
	expected := 4

	assert.Equal(t, expected, bulbs(A))
}

func TestSol2(t *testing.T) {
	A := []int{1}
	expected := 0

	assert.Equal(t, expected, bulbs(A))
}

func TestSol3(t *testing.T) {
	A := []int{0, 0, 0, 0}
	expected := 1

	assert.Equal(t, expected, bulbs(A))
}

func TestSol4(t *testing.T) {
	A := []int{1, 0, 0, 1, 1, 1}
	expected := 2

	assert.Equal(t, expected, bulbs(A))
}
