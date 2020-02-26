package solution

func isPalindrome(A string) int {
	n := len(A)
	if n <= 1 {
		return 1
	}
	l := 0
	h := n - 1
	for l < h {
		lch := convertChar(A[l])
		if lch == 0 {
			l++
			continue
		}
		hch := convertChar(A[h])
		if hch == 0 {
			h--
			continue
		}
		if lch != hch {
			return 0
		}
		l++
		h--
	}
	return 1
}

// convert upper case to lowwer case
// if char is not a letter return 0
func convertChar(ch byte) byte {
	if ch >= '0' && ch <= '9' {
		return ch
	}

	if ch >= 'a' && ch <= 'z' {
		return ch
	}

	if ch >= 'A' && ch <= 'Z' {
		return ch + 'a' - 'A'
	}

	return 0
}
