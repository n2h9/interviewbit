package solution

func solve(A string) int {
	n := len(A)
	if n <= 0 {
		return 0
	}

	mod := 10003
	sum := 0

	for i := 0; i < n; i++ {
		if isVowel(A[i]) {
			sum += n - i
		}
	}

	return sum % mod
}

func isVowel(ch byte) bool {
	vovel := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vovel {
		if ch == v {
			return true
		}
	}
	return false
}
