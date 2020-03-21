package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{1, 2, 3}
	expected := [][]int{
		[]int{},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 3},
		[]int{2},
		[]int{2, 3},
		[]int{3},
	}
	assert.Equal(t, expected, subsets(A))
}
