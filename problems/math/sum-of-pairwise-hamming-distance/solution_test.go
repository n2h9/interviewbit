package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{2, 7}
	expectedA := 4

	assert.Equal(t, expectedA, hammingDistance(testA))
}

func TestSol2(t *testing.T) {
	testA := []int{2, 4, 6}
	expectedA := 8

	assert.Equal(t, expectedA, hammingDistance(testA))
}

func TestSol3(t *testing.T) {
	testA := []int{1}
	expectedA := 0

	assert.Equal(t, expectedA, hammingDistance(testA))
}

func TestSol4(t *testing.T) {
	testA := []int{93, 73}
	expectedA := 4

	assert.Equal(t, expectedA, hammingDistance(testA))
}

func TestSol5(t *testing.T) {
	testA := []int{93, 73, 123, 123, 534, 23, 45, 11, 7890, 123123123, 45345343453, 8989892313, 123556546456, 222222, 11111111, 5}
	expectedA := 2596

	assert.Equal(t, expectedA, hammingDistance(testA))
}
