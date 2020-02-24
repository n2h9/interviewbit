package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 6
	expectedA := []int{1, 2, 3, 6}

	assert.Equal(t, expectedA, allFactors(testA))
}

func TestSol2(t *testing.T) {
	testA := 4
	expectedA := []int{1, 2, 4}

	assert.Equal(t, expectedA, allFactors(testA))
}

func TestSol3(t *testing.T) {
	testA := 11
	expectedA := []int{1, 11}

	assert.Equal(t, expectedA, allFactors(testA))
}

func TestSol4(t *testing.T) {
	testA := 24
	expectedA := []int{1, 2, 3, 4, 6, 8, 12, 24}

	assert.Equal(t, expectedA, allFactors(testA))
}

func TestSol5(t *testing.T) {
	testA := 25
	expectedA := []int{1, 5, 25}

	assert.Equal(t, expectedA, allFactors(testA))
}

func TestSol6(t *testing.T) {
	testA := 3
	expectedA := []int{1, 3}

	assert.Equal(t, expectedA, allFactors(testA))
}
