package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{1, 2, 2, 3, 1}
	expected := 3

	assert.Equal(t, expected, singleNumber(testA))
}

func TestSol1(t *testing.T) {
	testA := []int{1, 2, 2}
	expected := 1

	assert.Equal(t, expected, singleNumber(testA))
}
