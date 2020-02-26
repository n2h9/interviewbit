package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	hay := []int{23, 29, 31, 37, 41, 43, 47, 53, 59, 1, 3, 5, 7, 11, 13, 17, 19}
	needle := 31
	expectedIndex := 2

	assert.Equal(t, expectedIndex, search(hay, needle))
}

func TestSearch2(t *testing.T) {
	hay := []int{53, 59, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	needle := 42
	expectedIndex := -1

	assert.Equal(t, expectedIndex, search(hay, needle))
}

func TestSearch3(t *testing.T) {
	hay := []int{4, 5, 6, 7, 0, 1, 2, 3}
	needle := 4
	expectedIndex := 0

	assert.Equal(t, expectedIndex, search(hay, needle))
}

func TestSearch4(t *testing.T) {
	hay := []int{5, 17, 100, 3}
	needle := 6
	expectedIndex := -1

	assert.Equal(t, expectedIndex, search(hay, needle))
}
