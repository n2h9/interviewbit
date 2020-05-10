package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{2, 1, 2}
	expected := 2

	assert.Equal(t, expected, majorityElement(A))
}
