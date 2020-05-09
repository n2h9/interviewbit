package solution

import (
	"sort"
)

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func maxp3(A []int) int {
	if len(A) <= 3 {
		val := 1
		for i := 0; i < len(A); i++ {
			val *= A[i]
		}
		return val
	}

	sort.Ints(A)
	if A[0] < 0 && A[1] < 0 {
		v1 := A[0] * A[1] * A[len(A)-1]
		v2 := A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
		if v1 > v2 {
			return v1
		}
		return v2
	}
	return A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
}
