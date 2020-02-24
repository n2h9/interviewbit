package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 12121
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol2(t *testing.T) {
	testA := 123
	expectedA := 0

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol3(t *testing.T) {
	testA := -12121
	expectedA := 0

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol4(t *testing.T) {
	testA := 7
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol5(t *testing.T) {
	testA := 11
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol6(t *testing.T) {
	testA := 10
	expectedA := 0

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol7(t *testing.T) {
	testA := 111111111111
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol8(t *testing.T) {
	testA := 2147447412
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol9(t *testing.T) {
	testA := 1000021
	expectedA := 0

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol10(t *testing.T) {
	testA := 100001
	expectedA := 1

	assert.Equal(t, expectedA, isPalindrome(testA))
}

func TestSol11(t *testing.T) {
	testA := 111002111
	expectedA := 0

	assert.Equal(t, expectedA, isPalindrome(testA))
}
