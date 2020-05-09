package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{4, -4, 2}
	B := []int{4, 0, 5}
	expected := 4

	assert.Equal(t, expected, mice(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{-10, -9, 0, 3}
	B := []int{-15, -10, 3, 4}
	expected := 5

	assert.Equal(t, expected, mice(A, B))
}

func TestSol3(t *testing.T) {
	A := []int{-4, 2, 4}
	B := []int{-8, -7, -6, 0, 4, 5, 6, 7, 8}
	expected := 2

	assert.Equal(t, expected, mice(A, B))
}

func TestSol4(t *testing.T) {
	A := []int{-4, 0, 5, 12}
	B := []int{-10, -8, 0, 2, 6}
	expected := 6

	assert.Equal(t, expected, mice(A, B))
}

func TestSol5(t *testing.T) {
	A := []int{-4, 0, 5, 12}
	B := []int{-15, -12, 0, 2, 6}
	expected := 8

	assert.Equal(t, expected, mice(A, B))
}
