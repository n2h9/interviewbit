package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStraightSol(t *testing.T) {
	testA := []int{0, 2, 5, 7}
	expected := 2

	assert.Equal(t, expected, findMinXorStraight(testA))
}

func TestStraightSol1(t *testing.T) {
	testA := []int{0, 4, 7, 9}
	expected := 3

	assert.Equal(t, expected, findMinXorStraight(testA))
}
