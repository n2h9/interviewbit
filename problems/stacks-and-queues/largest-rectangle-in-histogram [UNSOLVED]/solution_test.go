package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{2, 1, 5, 6, 2, 3}
	expected := 10
	assert.Equal(t, expected, largestRectangleArea(A))
}

func TestSol2(t *testing.T) {
	A := []int{1}
	expected := 1
	assert.Equal(t, expected, largestRectangleArea(A))
}

func TestSol3(t *testing.T) {
	A := []int{90, 58, 69, 70, 82, 100, 13, 57, 47, 18}
	expected := 348
	assert.Equal(t, expected, largestRectangleArea(A))
}

func TestSol4(t *testing.T) {
	A := []int{6, 2, 5, 4, 5, 1, 6}
	expected := 12
	assert.Equal(t, expected, largestRectangleArea(A))
}

func TestSol5(t *testing.T) {
	A := []int{69, 47, 84, 7, 70, 73, 4, 73, 70, 54, 2, 35, 32, 53, 99, 41, 90, 53, 55, 6, 1, 95, 63, 63, 74, 12, 32, 89, 69, 71, 17, 49, 9, 40, 10, 56}
	expected := 256
	assert.Equal(t, expected, largestRectangleArea(A))
}
