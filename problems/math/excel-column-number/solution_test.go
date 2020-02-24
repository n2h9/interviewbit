package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := "CA"
	expectedA := 79

	assert.Equal(t, expectedA, titleToNumber(testA))
}

func TestSol2(t *testing.T) {
	testA := "A"
	expectedA := 1

	assert.Equal(t, expectedA, titleToNumber(testA))
}

func TestSol3(t *testing.T) {
	testA := "Z"
	expectedA := 26

	assert.Equal(t, expectedA, titleToNumber(testA))
}

func TestSol5(t *testing.T) {
	testA := "AAAAAAAAAAA"
	expectedA := 146813779479511

	assert.Equal(t, expectedA, titleToNumber(testA))
}

func TestSol6(t *testing.T) {
	testA := "AAA"
	expectedA := 703

	assert.Equal(t, expectedA, titleToNumber(testA))
}
