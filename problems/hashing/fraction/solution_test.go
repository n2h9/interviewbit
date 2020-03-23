package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := 1
	B := 2
	expected := "0.5"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol2(t *testing.T) {
	A := 2
	B := 1
	expected := "2"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol3(t *testing.T) {
	A := 2
	B := 3
	expected := "0.(6)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol4(t *testing.T) {
	A := 22
	B := 7
	expected := "3.(142857)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol5(t *testing.T) {
	A := 7
	B := 6
	expected := "1.1(6)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol6(t *testing.T) {
	A := 8
	B := 7
	expected := "1.(142857)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol7(t *testing.T) {
	A := -1
	B := 2
	expected := "-0.5"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol8(t *testing.T) {
	A := 8
	B := -4
	expected := "-2"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol9(t *testing.T) {
	A := -25
	B := -9
	expected := "2.(7)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}

func TestSol10(t *testing.T) {
	A := 945
	B := 103
	expected := "9.(1747572815533980582524271844660194)"

	assert.Equal(t, expected, fractionToDecimal(A, B))
}
