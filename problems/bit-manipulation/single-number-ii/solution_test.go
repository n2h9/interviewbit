package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}
	expected := 4

	assert.Equal(t, expected, singleNumber(testA))
}

func TestSol1(t *testing.T) {
	testA := []int{0, 0, 0, 1}
	expected := 1

	assert.Equal(t, expected, singleNumber(testA))
}

func TestSo2(t *testing.T) {
	testA := []int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1, 5, 5, 5}
	expected := 4

	assert.Equal(t, expected, singleNumber(testA))
}
