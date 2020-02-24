package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolB(t *testing.T) {
	testA := []int{1, 2, 0}
	expectedA := 3

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}

func TestSolB2(t *testing.T) {
	testA := []int{-8, 8, -9}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}

func TestSolB3(t *testing.T) {
	testA := []int{100, 100, 100}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}

func TestSolB4(t *testing.T) {
	testA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedA := 11

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}

func TestSolB5(t *testing.T) {
	testA := []int{}
	expectedA := 1

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}

func TestSolB6(t *testing.T) {
	testA := []int{4, 2, 10, 5, 3, 1}
	expectedA := 6

	assert.Equal(t, expectedA, firstMissingPositiveB(testA))
}
