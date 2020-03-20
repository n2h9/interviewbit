package solution

func Mod(A int, B int, C int) int {
	v := powMod(A, B, C)
	// this is for test case A = -1 B = 1 C = 20
	if v < 0 && C > 0 {
		return v + C
	}
	return v
}

func powMod(a, b, c int) int {
	if a == 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a % c
	}

	if isEven(b) {
		halfPow := powMod(a, b>>1, c)
		return halfPow * halfPow % c
	}

	return a % c * powMod(a, b-1, c) % c
}

func isEven(x int) bool {
	return x&1 == 0
}
