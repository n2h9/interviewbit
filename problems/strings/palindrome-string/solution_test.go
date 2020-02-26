package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := "A man, a plan, a canal: Panama"
	expected := 1

	assert.Equal(t, expected, isPalindrome(testA))
}

func TestSol2(t *testing.T) {
	testA := "race a car"
	expected := 0

	assert.Equal(t, expected, isPalindrome(testA))
}

func TestSol3(t *testing.T) {
	testA := "0123456789aaa   aa  a a a a  ,, , , ,aaa9876543210"
	expected := 1

	assert.Equal(t, expected, isPalindrome(testA))
}

func TestSol4(t *testing.T) {
	testA := "aaaaaaaaaaaaw"
	expected := 0

	assert.Equal(t, expected, isPalindrome(testA))
}

func TestSol5(t *testing.T) {
	testA := "*,****"
	expected := 1

	assert.Equal(t, expected, isPalindrome(testA))
}
