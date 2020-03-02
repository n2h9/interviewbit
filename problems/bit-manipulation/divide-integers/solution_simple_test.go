package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSol(t *testing.T) {
	A := 5
	B := 2
	expected := 2

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol2(t *testing.T) {
	A := 100
	B := 10
	expected := 10

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol3(t *testing.T) {
	A := 100
	B := 100000
	expected := 0

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol4(t *testing.T) {
	A := 100
	B := 0
	expected := 2147483647

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol5(t *testing.T) {
	A := -5700836
	B := -7169730

	expected := 0

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol6(t *testing.T) {
	A := -8
	B := 2

	expected := -4

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol7(t *testing.T) {
	A := 1
	B := -1

	expected := -1

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol8(t *testing.T) {
	A := 4
	B := 4

	expected := 1

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol9(t *testing.T) {
	A := -2147483648
	B := 1

	expected := -2147483648

	assert.Equal(t, expected, divideSimple(A, B))
}

func TestSimpleSol10(t *testing.T) {
	A := -2147483648
	B := -1

	expected := 2147483647

	assert.Equal(t, expected, divideSimple(A, B))
}
