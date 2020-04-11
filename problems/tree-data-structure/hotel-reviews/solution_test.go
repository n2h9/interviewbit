package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	A := "cool_ice_wifi"
	B := []string{"water_is_cool", "cold_ice_drink", "cool_wifi_speed"}
	expected := []int{2, 0, 1}

	assert.Equal(t, expected, solve(A, B))
}

func TestSol2(t *testing.T) {
	A := "test_sky_jail"
	B := []string{"test_rail_sky", "test_test_test_test", "sky_sky_jail"}
	expected := []int{1, 2, 0}

	assert.Equal(t, expected, solve(A, B))
}

func TestSol3(t *testing.T) {
	A := "test_sky_jail"
	B := []string{"test_sky", "sky_jail_test", "sky_sky_sky", "test_test"}
	expected := []int{1, 2, 0, 3}

	assert.Equal(t, expected, solve(A, B))
}
