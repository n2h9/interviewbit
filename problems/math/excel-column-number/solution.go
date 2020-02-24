package solution

func titleToNumber(A string) int {
	l := len(A)
	if l <= 0 {
		return 0
	}
	const base = 26

	n := 0
	for i := l - 1; i >= 0; i-- {
		n += charToInt(A[i]) * power(base, l-1-i)
	}

	return n
}

func charToInt(b byte) int {
	// 64 = ASCII before A; so A = 1, B = 2 .. Z = 26
	return int(b - 64)
}

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	r := a
	for ; b > 1; b-- {
		r *= a
	}
	return r
}
