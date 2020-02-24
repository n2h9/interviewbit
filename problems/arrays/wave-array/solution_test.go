package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{1, 2, 3, 4}
	expectedA := []int{2, 1, 4, 3}

	assert.Equal(t, expectedA, wave(testA))
}

func TestSol1(t *testing.T) {
	testA := []int{1}
	expectedA := []int{1}

	assert.Equal(t, expectedA, wave(testA))
}

func TestSo2(t *testing.T) {
	testA := []int{100, 27, 19, 1, 4}
	expectedA := []int{4, 1, 27, 19, 100}

	assert.Equal(t, expectedA, wave(testA))
}

func TestSol3(t *testing.T) {
	testA := []int{200, 1000, 2, 18, 5, 20}
	expectedA := []int{5, 2, 20, 18, 1000, 200}

	assert.Equal(t, expectedA, wave(testA))
}
