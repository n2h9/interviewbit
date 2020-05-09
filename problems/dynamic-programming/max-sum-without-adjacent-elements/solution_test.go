package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := [][]int{[]int{1}, []int{2}}
	expected := 2

	assert.Equal(t, expected, adjacent(A))
}

func TestSol1(t *testing.T) {
	A := [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 4, 5},
	}
	expected := 8

	assert.Equal(t, expected, adjacent(A))
}

func TestSol2(t *testing.T) {
	// понятно
	// 31+70+82+77+61
	A := [][]int{
		[]int{16, 5, 54, 55, 36, 82, 61, 77, 66, 61},
		[]int{31, 30, 36, 70, 9, 37, 1, 11, 68, 14},
	}
	expected := 321

	assert.Equal(t, expected, adjacent(A))
}
