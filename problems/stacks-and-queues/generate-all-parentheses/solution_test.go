package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	assert.Equal(t, 1, isValid(""))
	assert.Equal(t, 1, isValid("()"))
	assert.Equal(t, 1, isValid("()[]{}"))
}

func TestSol2(t *testing.T) {
	assert.Equal(t, 0, isValid("(]"))
	assert.Equal(t, 0, isValid("([)]"))
	assert.Equal(t, 0, isValid("[{"))
}
