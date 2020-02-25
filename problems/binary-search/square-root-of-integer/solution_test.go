package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	testA := 4
	expected := 2

	assert.Equal(t, expected, sqrt(testA))
}

func TestSol2(t *testing.T) {
	testA := 16
	expected := 4

	assert.Equal(t, expected, sqrt(testA))
}

func TestSol3(t *testing.T) {
	testA := 900
	expected := 30

	assert.Equal(t, expected, sqrt(testA))
}

func TestSol4(t *testing.T) {
	testA := 15876
	expected := 126

	assert.Equal(t, expected, sqrt(testA))
}

func TestSol5(t *testing.T) {
	testA := 15875
	expected := 125

	assert.Equal(t, expected, sqrt(testA))
}

func TestSol6(t *testing.T) {
	assert.Equal(t, 0, sqrt(0))
	assert.Equal(t, 1, sqrt(3))
	assert.Equal(t, 2, sqrt(4))
}
