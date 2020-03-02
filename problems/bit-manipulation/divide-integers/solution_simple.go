package solution

func divideSimple(A int, B int) int {
	// division by zero is not allowed :)
	if B == 0 {
		return MAX_INT
	}
	var shouldChangeResSign bool
	switch {
	case A < 0 && B > 0:
		A = oppositeNumSimple(A)
		shouldChangeResSign = true
	case A > 0 && B < 0:
		B = oppositeNumSimple(B)
		shouldChangeResSign = true
	case A < 0 && B < 0:
		A = oppositeNumSimple(A)
		B = oppositeNumSimple(B)
	}
	if A < B {
		return 0
	}

	val := 0
	switch {
	case B == 1:
		val = A
	case A == B:
		val = 1
	default:
		for ; A >= B; A -= B {
			val++
		}
	}

	if shouldChangeResSign {
		val = oppositeNumSimple(val)
	}
	// overflow case
	if val > MAX_INT || val < MIN_INT {
		return MAX_INT
	}
	return val
}

func oppositeNumSimple(x int) int {
	return (^x) + 1
}
