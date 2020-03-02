package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 5
	B := 2
	expected := 2

	assert.Equal(t, expected, divide(A, B))
}

func TestSol2(t *testing.T) {
	A := 100
	B := 10
	expected := 10

	assert.Equal(t, expected, divide(A, B))
}

func TestSol3(t *testing.T) {
	A := 100
	B := 100000
	expected := 0

	assert.Equal(t, expected, divide(A, B))
}

func TestSol4(t *testing.T) {
	A := 100
	B := 0
	expected := 2147483647

	assert.Equal(t, expected, divide(A, B))
}

func TestSol5(t *testing.T) {
	A := -5700836
	B := -7169730

	expected := 0

	assert.Equal(t, expected, divide(A, B))
}

func TestSol6(t *testing.T) {
	A := -8
	B := 2

	expected := -4

	assert.Equal(t, expected, divide(A, B))
}

func TestSol7(t *testing.T) {
	A := 1
	B := -1

	expected := -1

	assert.Equal(t, expected, divide(A, B))
}

func TestSol8(t *testing.T) {
	A := 4
	B := 4

	expected := 1

	assert.Equal(t, expected, divide(A, B))
}

func TestSol9(t *testing.T) {
	A := -2147483648
	B := 1

	expected := -2147483648

	assert.Equal(t, expected, divide(A, B))
}

func TestSol10(t *testing.T) {
	A := -2147483648
	B := -1

	expected := 2147483647

	assert.Equal(t, expected, divide(A, B))
}

func TestDiv(t *testing.T) {
	var (
		A    uint = 23456
		B    uint = 12
		expV uint = 1954
		expC uint = 8
	)

	v, c := div(A, B)
	assert.Equal(t, expV, v)
	assert.Equal(t, expC, c)
}
