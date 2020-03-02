package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := "ABEC"
	expected := 6

	assert.Equal(t, expected, solve(testA))
}
