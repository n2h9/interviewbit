package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve0(t *testing.T) {
	numRows := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	assert.Equal(t, expected, solve(numRows))
}

func TestSolve1(t *testing.T) {
	numRows := -1
	var expected [][]int
	assert.Equal(t, expected, solve(numRows))
}

func TestSolve2(t *testing.T) {
	numRows := 0
	var expected [][]int
	assert.Equal(t, expected, solve(numRows))
}

func TestSolve3(t *testing.T) {
	numRows := 1
	expected := [][]int{
		{1},
	}
	assert.Equal(t, expected, solve(numRows))
}

func TestSolve4(t *testing.T) {
	numRows := 6
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
		{1, 5, 10, 10, 5, 1},
	}
	assert.Equal(t, expected, solve(numRows))
}
