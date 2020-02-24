package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{1, 2, 0}
	expectedA := 3

	assert.Equal(t, expectedA, firstMissingPositive(testA))
}

func TestSol2(t *testing.T) {
	testA := []int{-8, 8, -9}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositive(testA))
}

func TestSol3(t *testing.T) {
	testA := []int{100, 100, 100}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositive(testA))
}

func TestSol4(t *testing.T) {
	testA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedA := 11

	assert.Equal(t, expectedA, firstMissingPositive(testA))
}

func TestSol5(t *testing.T) {
	testA := []int{}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositive(testA))
}
