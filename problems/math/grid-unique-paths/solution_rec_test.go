package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolRec(t *testing.T) {
	testA := 2
	testB := 2
	expectedA := 2

	assert.Equal(t, expectedA, uniquePathsRec(testA, testB))
}

func TestSolRec2(t *testing.T) {
	testA := 1000
	testB := 1
	expectedA := 1

	assert.Equal(t, expectedA, uniquePathsRec(testA, testB))
}

func TestSolRec3(t *testing.T) {
	testA := 7
	testB := 3
	expectedA := 28

	assert.Equal(t, expectedA, uniquePathsRec(testA, testB))
}

func TestSolRec4(t *testing.T) {
	testA := 3
	testB := 7
	expectedA := 28

	assert.Equal(t, expectedA, uniquePathsRec(testA, testB))
}
