package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve0(t *testing.T) {
	A := []int{3, 4, 1, 4, 1}
	expected := 4
	assert.Equal(t, expected, repeatedNumber(A))
}

func TestSolve1(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	expected := 1
	assert.Equal(t, expected, repeatedNumber(A))
}

func TestSolve2(t *testing.T) {
	A := []int{1, 2, 2}
	expected := 2
	assert.Equal(t, expected, repeatedNumber(A))
}

func TestSolve3(t *testing.T) {
	A := []int{1}
	expected := 1
	assert.Equal(t, expected, repeatedNumber(A))
}

func TestSolve4(t *testing.T) {
	A := []int{1, 1}
	expected := 1
	assert.Equal(t, expected, repeatedNumber(A))
}
