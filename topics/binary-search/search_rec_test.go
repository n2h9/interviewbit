package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRec(t *testing.T) {
	hay := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	needle := 37
	expectedIndex := 11

	assert.Equal(t, expectedIndex, searchRec(hay, needle, 0, len(hay)-1))
}

func TestSearchRec2(t *testing.T) {
	hay := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	needle := 42
	expectedIndex := -1

	assert.Equal(t, expectedIndex, searchRec(hay, needle, 0, len(hay)-1))
}
