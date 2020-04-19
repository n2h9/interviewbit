package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testA := []int{25, 17, 31, 13, 2}
	expectedA := []int{2, 13, 17, 25, 31}

	sort(testA)
	assert.Equal(t, expectedA, testA)
}

func TestSort2(t *testing.T) {
	testA := []int{99, 90, 88, 80, 77, 70, 66, 60, 55, 50, 44, 40, 33, 30}
	expectedA := []int{30, 33, 40, 44, 50, 55, 60, 66, 70, 77, 80, 88, 90, 99}

	sort(testA)
	assert.Equal(t, expectedA, testA)
}

func TestSort3(t *testing.T) {
	testA := []int{100, 100, 100, 100, 100}
	expectedA := []int{100, 100, 100, 100, 100}

	sort(testA)
	assert.Equal(t, expectedA, testA)
}

func TestSort4(t *testing.T) {
	testA := []int{10, 9}
	expectedA := []int{9, 10}

	sort(testA)
	assert.Equal(t, expectedA, testA)
}

func TestSort5(t *testing.T) {
	testA := []int{11, 2, 9, 13, 57, 25, 17, 1, 90, 3.}
	expectedA := []int{1, 2, 3, 9, 11, 13, 17, 25, 57, 90}

	sort(testA)
	assert.Equal(t, expectedA, testA)
}
