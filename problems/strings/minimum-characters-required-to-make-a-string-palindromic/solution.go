package solution

func solve(A string) int {
	n := len(A)
	if n <= 1 {
		return 0
	}
	if isPal(A, 0, n-1) {
		return 0
	}
	for i := 1; i < n; i++ {
		// if isPal(A, 0, n-1-i) || isPal(A, i, n-1) {
		// 	return i
		// }
		if isPal(A, 0, n-1-i) {
			return i
		}
	}

	return n - 1
}

func isPal(s string, start, end int) bool {
	border := start + (end-start)/2
	for i := start; i <= border; i++ {
		if s[i] != s[end-i+start] {
			return false
		}
	}
	return true
}
