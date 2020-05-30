package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := "hit"
	B := "cog"
	C := []string{"hot", "dot", "dog", "lot", "log"}
	expected := 5

	assert.Equal(t, expected, solve(A, B, C))
}
