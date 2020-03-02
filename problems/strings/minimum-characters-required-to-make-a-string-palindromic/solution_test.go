package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := "ABC"
	expected := 2

	assert.Equal(t, expected, solve(testA))
}

func TestSol2(t *testing.T) {
	testA := "AACECAAAA"
	expected := 2

	assert.Equal(t, expected, solve(testA))
}

func TestSol3(t *testing.T) {
	testA := "ABCD"
	expected := 3

	assert.Equal(t, expected, solve(testA))
}

func TestSol4(t *testing.T) {
	testA := "AAAAAAAAAAAAAAAhhhhHYWTEASDOASUDASDJKHASJKDHASJKDHASJKDASHDJKASDHASKDHASDJKAHDJKHQWEUIQWEDIHASJKDHASJKDHASDJKASHDJKASDASDASDASD"
	expected := 112

	assert.Equal(t, expected, solve(testA))
}

func TestSol5(t *testing.T) {
	testA := "AAA"
	expected := 0

	assert.Equal(t, expected, solve(testA))
}

func TestSol6(t *testing.T) {
	testA := "banana"
	expected := 5

	assert.Equal(t, expected, solve(testA))
}
