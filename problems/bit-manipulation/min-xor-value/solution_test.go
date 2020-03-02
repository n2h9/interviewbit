package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{0, 2, 5, 7}
	expected := 2

	assert.Equal(t, expected, findMinXor(testA))
}

func TestSol1(t *testing.T) {
	testA := []int{0, 4, 7, 9}
	expected := 3

	assert.Equal(t, expected, findMinXor(testA))
}

func TestSol2(t *testing.T) {
	testA := []int{0, 4, 7, 9, 11, 23, 17, 90, 91, 92, 93, 123123, 12312314444, 1231231231, 12313131}
	expected := 1

	assert.Equal(t, expected, findMinXor(testA))
}
