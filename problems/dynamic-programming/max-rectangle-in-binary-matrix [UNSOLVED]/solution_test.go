package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSol(t *testing.T) {
// 	A := [][]int{
// 		[]int{1, 1, 1},
// 		[]int{0, 1, 1},
// 		[]int{1, 0, 0},
// 	}
// 	expected := 4

// 	assert.Equal(t, expected, maximalRectangle(A))
// }

// func TestSol1(t *testing.T) {
// 	A := [][]int{
// 		[]int{1, 1, 1},
// 		[]int{1, 1, 1},
// 		[]int{1, 1, 1},
// 	}
// 	expected := 9

// 	assert.Equal(t, expected, maximalRectangle(A))
// }

// func TestSol2(t *testing.T) {
// 	A := [][]int{
// 		[]int{1, 1, 0},
// 		[]int{1, 0, 1},
// 		[]int{1, 1, 1},
// 	}
// 	expected := 3

// 	assert.Equal(t, expected, maximalRectangle(A))
// }

// func TestSol3(t *testing.T) {
// 	A := [][]int{
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 		[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
// 	}
// 	expected := 100

// 	assert.Equal(t, expected, maximalRectangle(A))
// }

func TestSol4(t *testing.T) {
	A := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 0},
		[]int{1, 1, 1, 1, 1, 1, 1, 0},
		[]int{1, 1, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 1, 1, 0, 0, 0},
	}
	// :(
	expected := 21

	assert.Equal(t, expected, maximalRectangle(A))
}
