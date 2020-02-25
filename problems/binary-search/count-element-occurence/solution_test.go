package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := []int{5, 7, 7, 8, 8, 8, 8, 8, 10}
	b := 8
	expected := 5

	assert.Equal(t, expected, findCount(testA, b))
}

func TestSol2(t *testing.T) {
	testA := []int{8, 8, 8, 8, 8, 8, 8, 8, 8}
	b := 8
	expected := 9

	assert.Equal(t, expected, findCount(testA, b))
}

func TestSol3(t *testing.T) {
	testA := []int{}
	b := 8
	expected := 0

	assert.Equal(t, expected, findCount(testA, b))
}

func TestSol4(t *testing.T) {
	testA := []int{8}
	b := 8
	expected := 1

	assert.Equal(t, expected, findCount(testA, b))
}
