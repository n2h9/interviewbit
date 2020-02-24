package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 2
	expectedA := [][]int{
		{2, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
	}

	assert.Equal(t, expectedA, prettyPrint(testA))
}

func TestSol2(t *testing.T) {
	testA := 3
	expectedA := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}

	assert.Equal(t, expectedA, prettyPrint(testA))
}
