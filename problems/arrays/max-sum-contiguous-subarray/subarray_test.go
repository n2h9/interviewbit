package subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray0(t *testing.T) {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	assert.Equal(t, 6, maxSubArray(a))
}

func TestMaxSubArray1(t *testing.T) {
	a := []int{1, 2, 3, 4, -10}
	assert.Equal(t, 10, maxSubArray(a))
}

func TestMaxSubArray2(t *testing.T) {
	a := []int{0, 0, 0, 0, 0}
	assert.Equal(t, 0, maxSubArray(a))
}

func TestMaxSubArray3(t *testing.T) {
	a := []int{-1000, -1000, -1000, -1000, -1000}
	assert.Equal(t, -1000, maxSubArray(a))
}

func TestMaxSubArray4(t *testing.T) {
	a := []int{1000, -1000, -1000, -1000, -1000}
	assert.Equal(t, 1000, maxSubArray(a))
}

func TestMaxSubArray5(t *testing.T) {
	a := []int{1000, -1000, -1000, -1000, 10000}
	assert.Equal(t, 10000, maxSubArray(a))
}

func TestMaxSubArray6(t *testing.T) {
	a := []int{99, -1000, 100, -1000, 10000}
	assert.Equal(t, 10000, maxSubArray(a))
}

func TestMaxSubArray7(t *testing.T) {
	a := []int{600, 500, -1000, 1000, 150}
	assert.Equal(t, 1250, maxSubArray(a))
}
