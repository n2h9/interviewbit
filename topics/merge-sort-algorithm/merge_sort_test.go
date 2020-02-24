package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testA := []int{25, 17, 31, 13, 2, 1, 0, 11, 10, 9}
	expectedA := []int{0, 1, 2, 9, 10, 11, 13, 17, 25, 31}

	sort(testA, 0, len(testA)-1)
	assert.Equal(t, expectedA, testA)
}

func TestSort2(t *testing.T) {
	testA := []int{99, 90, 88, 80, 77, 70, 66, 60, 55, 50, 44, 40, 33, 30}
	expectedA := []int{30, 33, 40, 44, 50, 55, 60, 66, 70, 77, 80, 88, 90, 99}

	sort(testA, 0, len(testA)-1)
	assert.Equal(t, expectedA, testA)
}

func TestSort3(t *testing.T) {
	testA := []int{100, 100, 100, 100, 100}
	expectedA := []int{100, 100, 100, 100, 100}

	sort(testA, 0, len(testA)-1)
	assert.Equal(t, expectedA, testA)
}
