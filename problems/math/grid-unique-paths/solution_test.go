package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 2
	testB := 2
	expectedA := 2

	assert.Equal(t, expectedA, uniquePaths(testA, testB))
}

func TestSol2(t *testing.T) {
	testA := 1000
	testB := 1
	expectedA := 1

	assert.Equal(t, expectedA, uniquePaths(testA, testB))
}

func TestSol3(t *testing.T) {
	testA := 7
	testB := 3
	expectedA := 28

	assert.Equal(t, expectedA, uniquePaths(testA, testB))
}

func TestSol4(t *testing.T) {
	testA := 7
	testB := 4
	expectedA := 84

	assert.Equal(t, expectedA, uniquePaths(testA, testB))
}

func TestSol5(t *testing.T) {
	testA := 10
	testB := 10
	expectedA := 48620

	assert.Equal(t, expectedA, uniquePaths(testA, testB))
}
