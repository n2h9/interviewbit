package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{1, 3, 5}
	expected := 8

	assert.Equal(t, expected, cntBits(testA))
}

func TestSol2(t *testing.T) {
	testA := []int{1, 3, 5, 12, 11, 45, 100, 10000000, 12312313}
	expected := 528

	assert.Equal(t, expected, cntBits(testA))
}
