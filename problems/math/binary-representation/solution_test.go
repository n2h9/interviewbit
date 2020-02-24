package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 6
	expectedA := "110"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}

func TestSol2(t *testing.T) {
	testA := 0
	expectedA := "0"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}

func TestSol3(t *testing.T) {
	testA := 1
	expectedA := "1"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}

func TestSol4(t *testing.T) {
	testA := 2
	expectedA := "10"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}

func TestSol5(t *testing.T) {
	testA := 256
	expectedA := "100000000"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}

func TestSol6(t *testing.T) {
	testA := 255
	expectedA := "11111111"

	assert.Equal(t, expectedA, findDigitsInBinary(testA))
}
