package solution

import (
	"sort"
)

const max int = 1000000001

func findMinXor(A []int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return A[0]
	}

	sort.Ints(A)

	minXor := max
	var xor, val1, val2 int
	var i uint
	for i = 0; i < 32 && minXor == max; i++ {
		for j := 1; j < n; j++ {
			val1 = (A[j] >> i) << i
			val2 = (A[j-1] >> i) << i
			if val1 == val2 {
				xor = A[j] ^ A[j-1]
				if xor < minXor {
					minXor = xor
				}
			}
		}
	}

	return minXor
}
