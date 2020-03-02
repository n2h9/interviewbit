package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := []int{1, 3, 5}
	B := 4
	expected := 1

	assert.Equal(t, expected, diffPossible(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{1, 3, 5, 12, 15, 20, 26, 100, 200, 1000}
	B := 180
	expected := 1

	assert.Equal(t, expected, diffPossible(A, B))
}

func TestSol3(t *testing.T) {
	A := []int{1, 3, 4}
	B := 100
	expected := 0

	assert.Equal(t, expected, diffPossible(A, B))
}

func TestSol4(t *testing.T) {
	A := []int{1, 100}
	B := 99
	expected := 1

	assert.Equal(t, expected, diffPossible(A, B))
}

func TestSol5(t *testing.T) {
	A := []int{10, 20, 30, 40, 50, 60, 61, 62}
	B := 2
	expected := 1

	assert.Equal(t, expected, diffPossible(A, B))
}
