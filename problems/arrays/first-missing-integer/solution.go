package solution

// firstMissingPositive
// Not constant space
func firstMissingPositive(A []int) int {
	n := len(A)
	if n <= 0 {
		return 1
	}

	b := make([]bool, n)
	for i := 0; i < n; i++ {
		if A[i] <= n && A[i] > 0 {
			b[A[i]-1] = true
		}
	}

	for i := 0; i < n; i++ {
		if !b[i] {
			return i + 1
		}
	}

	return n + 1
}
