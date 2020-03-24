package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	assert.Equal(t, expected, longestConsecutive(A))
}

func TestSol2(t *testing.T) {
	A := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 89}
	expected := 10
	assert.Equal(t, expected, longestConsecutive(A))
}
