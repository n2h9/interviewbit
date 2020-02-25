package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{5, 7, 7, 8, 8, 8, 8, 8, 10}
	b := 8
	expected := []int{3, 7}

	assert.Equal(t, expected, searchRange(testA, b))
}

func TestSol2(t *testing.T) {
	testA := []int{8, 8, 8, 8, 8, 8, 8, 8, 8}
	b := 8
	expected := []int{0, 8}

	assert.Equal(t, expected, searchRange(testA, b))
}

func TestSol3(t *testing.T) {
	testA := []int{}
	b := 8
	expected := []int{-1, -1}

	assert.Equal(t, expected, searchRange(testA, b))
}

func TestSol4(t *testing.T) {
	testA := []int{8}
	b := 8
	expected := []int{0, 0}

	assert.Equal(t, expected, searchRange(testA, b))
}
