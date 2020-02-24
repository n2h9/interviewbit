package spiral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @input A : 2D integer array
 *
 * @Output Integer array.
 */
func TestSpiralOrder1(t *testing.T) {
	A := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	assert.Equal(t, expected, spiralOrder(A))
}

func TestSpiralOrder2(t *testing.T) {
	A := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
	}
	expected := []int{1, 2, 3, 4, 5, 10, 15, 14, 13, 12, 11, 6, 7, 8, 9}
	assert.Equal(t, expected, spiralOrder(A))
}

func TestSpiralOrder3(t *testing.T) {
	A := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{10, 11, 12},
		[]int{13, 14, 15},
	}
	expected := []int{1, 2, 3, 6, 9, 12, 15, 14, 13, 10, 7, 4, 5, 8, 11}
	assert.Equal(t, expected, spiralOrder(A))
}
