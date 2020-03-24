package solution

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func longestConsecutive(A []int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}

	V := make(map[int]int, n)
	maxLen := 1

	for i := 0; i < n; i++ {
		V[A[i]] = 1
	}

	for i := 0; i < n; i++ {
		if V[A[i]] != 1 {
			// already processed
			continue
		}
		v := calculateLength(V, A[i])
		if v > maxLen {
			maxLen = v
		}
	}

	return maxLen
}

// calculate length of consecutive sequence
// save to map
// and return as value
func calculateLength(V map[int]int, index int) int {
	v := 1
	if V[index-1] > 0 {
		v = 1 + calculateLength(V, index-1)
	}
	V[index] = v
	return v
}
