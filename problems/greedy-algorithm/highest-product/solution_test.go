package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{1, 2, 3, 4}
	expected := 24

	assert.Equal(t, expected, maxp3(A))
}

func TestSol1(t *testing.T) {
	A := []int{0, -1, 3, 100, 70, 50}
	expected := 350000

	assert.Equal(t, expected, maxp3(A))
}

func TestSol2(t *testing.T) {
	A := []int{0, -1, -10000, 3, 100, 70, 50}
	expected := 1000000

	assert.Equal(t, expected, maxp3(A))
}

func TestSol3(t *testing.T) {
	A := []int{-3, -1, 0, 0, 1, 4}
	expected := 12

	assert.Equal(t, expected, maxp3(A))
}

func TestSol4(t *testing.T) {
	A := []int{-1, -1, -1, -1, -1, -1}
	expected := -1

	assert.Equal(t, expected, maxp3(A))
}

func TestSol5(t *testing.T) {
	A := []int{100, 100, 100, 100, 100, 100}
	expected := 1000000

	assert.Equal(t, expected, maxp3(A))
}

func TestSol6(t *testing.T) {
	A := []int{-1, 2, 3}
	expected := -6

	assert.Equal(t, expected, maxp3(A))
}
