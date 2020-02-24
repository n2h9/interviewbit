package zeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve0(t *testing.T) {
	A := []int{3, 4, 1, 4, 1}
	expected := 4
	assert.Equal(t, expected, repeatedNumber(A))
}
